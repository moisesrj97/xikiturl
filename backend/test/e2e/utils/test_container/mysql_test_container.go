package test_container

import (
	"crypto/rand"
	"time"
)

func MySqlTestContainer() TestContainer {
	info := ContainerInfo{
		Image: "mysql",
		PortMappings: []PortMapping{
			{HostPort: "3306", ContainerPort: "3306"},
		},
		Environment: []Environment{
			{Key: "MYSQL_ROOT_PASSWORD", Value: "password"},
			{Key: "MYSQL_DATABASE", Value: "xikiturl"},
		},
		WaitLog:      "/usr/sbin/mysqld: ready for connections",
		StartTimeout: 30 * time.Second,
	}
	name := "test-container-" + rand.Text()
	testContainer := TestContainer{
		Name: name,
		Info: info,
	}
	return testContainer
}
