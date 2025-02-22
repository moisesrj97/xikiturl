package test_container

import (
	"fmt"
)

type Args []string

func (a *Args) AddImageName(name string) {
	*a = append(*a, name)
}

func (a *Args) AddEnv(envs []Environment) {
	for _, env := range envs {
		*a = append(*a, "-e", fmt.Sprintf("%s=%s", env.Key, env.Value))
	}
}

func (a *Args) AddPortMappings(ports []int) {
	for _, port := range ports {
		*a = append(*a, "-p", fmt.Sprintf("%d", port))
	}
}
