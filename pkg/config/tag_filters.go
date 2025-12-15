package config

import (
	"regexp"
	"strings"
)

type TagFilters struct {
	filters map[string]string
}

func NewTagFilters(tfs []string) *TagFilters {
	tagFilters := map[string]string{}

	for _, tf := range tfs {
		parts := strings.Split(tf, "=")
		tagFilters[parts[0]] = parts[1]
	}

	return &TagFilters{
		filters: tagFilters,
	}
}

func (t *TagFilters) IsMatch(gitRepo *GitRepo) (bool, error) {
	for tag, value := range t.filters {
		repoVal, tagFound := gitRepo.Tags[tag]
		if !tagFound {
			return false, nil
		}

		match, err := regexp.Match(value, []byte(repoVal))
		if err != nil {
			return false, err
		}

		if !match { 
			return false, nil
		}
	}

	return true, nil
}