# gitbatch

A cli to perform batch operations against git repositories.

## Configuration

The global configuration file is place in 

- %userprofile%/.gitbatch/config
- ~/.gitbatch/config

You can print the config

    gobatch config print

Or get the path to the config

    gobatch config path

## Exec shell commands

    gobatch exec -c "ls"

    gobatch exec -c "ls" -t wsl