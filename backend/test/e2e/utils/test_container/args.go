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

func (a *Args) AddPortMappings(portMappings []PortMapping) {
	for _, port := range portMappings {
		*a = append(*a, "-p", fmt.Sprintf("%s:%s", port.HostPort, port.ContainerPort))
	}
}
