package test_container

import (
	"crypto/rand"
	"fmt"
	"time"
)

func MySqlTestContainer() (TestContainer, string) {
	password := "password"
	dbName := "xikiturl"
	port := "3306"

	info := ContainerInfo{
		Image: "mysql",
		PortMappings: []PortMapping{
			{HostPort: port, ContainerPort: port},
		},
		Environment: []Environment{
			{Key: "MYSQL_ROOT_PASSWORD", Value: password},
			{Key: "MYSQL_DATABASE", Value: dbName},
		},
		WaitLog:      "X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /var/run/mysqld/mysqlx.sock",
		StartTimeout: 30 * time.Second,
	}
	name := "test-container-" + rand.Text()
	testContainer := TestContainer{
		Name: name,
		Info: info,
	}
	return testContainer, fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"root",
		password,
		"127.0.0.1:"+port,
		dbName,
	)
}
