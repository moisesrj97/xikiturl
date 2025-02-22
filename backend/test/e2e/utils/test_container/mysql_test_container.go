package test_container

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"time"
)

var password = "password"
var dbName = "xikiturl"

type MySqlTestContainer struct {
	TestContainer
}

func NewMySqlTestContainer() MySqlTestContainer {
	info := ContainerInfo{
		Image: "mysql",
		Ports: []int{3306},
		Environment: []Environment{
			{Key: "MYSQL_ROOT_PASSWORD", Value: password},
			{Key: "MYSQL_DATABASE", Value: dbName},
		},
		WaitLog:      "X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /var/run/mysqld/mysqlx.sock",
		StartTimeout: 30 * time.Second,
	}
	name := "test-container-" + rand.Text()
	testContainer := MySqlTestContainer{
		TestContainer: TestContainer{
			Name: name,
			Info: info,
		},
	}
	return testContainer
}

func (mq *MySqlTestContainer) Start() (string, error) {
	portMappings, err := mq.TestContainer.Start()

	if err != nil {
		return "", err
	}

	return generateConnectionString(portMappings), nil
}

func generateConnectionString(portMappings []PortMapping) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"root",
		password,
		"127.0.0.1:"+strconv.Itoa(portMappings[0].HostPort),
		dbName,
	)
}
