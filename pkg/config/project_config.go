package config

import (
	"errors"
	"io/fs"
	"os"
	"path"

	"github.com/Useurmind/gitbatch/pkg/output"
	"gopkg.in/yaml.v2"
)

type ProjectConfig struct {
	GitRepos []*GitRepo `yaml:"gitRepos,omitempty"`
}

func (c *ProjectConfig) GetGitRepos() []*GitRepo {
	if c.GitRepos == nil {
		c.GitRepos = make([]*GitRepo, 0)
	}

	return c.GitRepos
}

func (c *ProjectConfig) HasGitRepo(path string) bool {
	for _, gitrepo := range c.GetGitRepos() {
		if gitrepo.Path == path {
			return true
		}
	}

	return false
}

type GitRepo struct {
	Name string `yaml:"name,omitempty"`
	Path string `yaml:"path,omitempty"`
	Tags map[string]string `yaml:"tags,omitempty"`
}

func (c *GitRepo) GetName() string {
	if c.Name != "" {
		return c.Name
	}

	return c.Path
}

func GetProjectConfig(projectConfigPath string) (*ProjectConfig, error) {
	_, err :=  os.Stat(projectConfigPath)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	data, err := os.ReadFile(projectConfigPath)
	if err != nil {
		return nil, err
	}

	config := &ProjectConfig{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func WriteProjectConfig(projectConfigPath string, projectConfig *ProjectConfig) error {
	data, err :=yaml.Marshal(projectConfig)
	if err != nil {
		return err
	}

	return os.WriteFile(projectConfigPath, data, 0777)
}

func ExtendProjectFromSubdirOf(searchPath string, projectConfigPath string) error {
	projectConfig, err := GetProjectConfig(projectConfigPath)
	if err != nil {
		return err
	}
	if projectConfig == nil {
		projectConfig = &ProjectConfig{}
	}

	subDirs, err := os.ReadDir(searchPath)
	if err != nil {
		return err
	}
	for _, subdir := range subDirs {
		if !subdir.IsDir() {
			continue
		}
		gitRepoDir := path.Join(searchPath, subdir.Name())
		gitDir := path.Join(gitRepoDir, ".git")
		_, err=  os.Stat(gitDir)
		if errors.Is(err, fs.ErrNotExist) {
			output.Writeln("directory %s is not a git repo, ignore", gitRepoDir)
			continue
		}

		if !projectConfig.HasGitRepo(gitRepoDir) {
			projectConfig.GitRepos = append(projectConfig.GitRepos, &GitRepo{
				Name: subdir.Name(),
				Path: gitRepoDir,
				Tags: map[string]string{
					"name": subdir.Name(),
				},
			})
		}
	}

	return WriteProjectConfig(projectConfigPath, projectConfig)
} 
