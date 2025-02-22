package test_container

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type PortMapping struct {
	HostPort      int
	ContainerPort int
}

type Environment struct {
	Key   string
	Value string
}

type ContainerInfo struct {
	Image        string
	Ports        []int
	Environment  []Environment
	WaitLog      string
	StartTimeout time.Duration
}

var Debug = false

type TestContainer struct {
	Name string
	Info ContainerInfo
}

func (tc *TestContainer) Start() ([]PortMapping, error) {
	args := Args{"run", "--name", tc.Name}

	args.AddEnv(tc.Info.Environment)
	args.AddPortMappings(tc.Info.Ports)
	args.AddImageName(tc.Info.Image)

	command := exec.Command("docker", args...)

	pipe, err := command.StdoutPipe()
	if err != nil {
		return nil, err
	}

	pipeErr, err := command.StderrPipe()
	if err != nil {
		return nil, err
	}

	err = command.Start()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(pipe)
	scannerErr := bufio.NewScanner(pipeErr)

	ready := make(chan bool)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			if Debug {
				log.Println(line)
			}
			if strings.Contains(line, tc.Info.WaitLog) {
				ready <- true
			}
		}
	}()

	go func() {
		for scannerErr.Scan() {
			line := scannerErr.Text()
			if Debug {
				log.Println(line)
			}
			if strings.Contains(line, tc.Info.WaitLog) {
				ready <- true
			}
		}
	}()

	select {
	case <-ready:
	case <-time.After(tc.Info.StartTimeout):
		return nil, fmt.Errorf("Timed out waiting for container to start")
	}

	log.Printf("Started container: %s", tc.Name)

	var ports []PortMapping

	for _, containerPort := range tc.Info.Ports {
		cmd := exec.Command("docker", "port", tc.Name, strconv.Itoa(containerPort))
		output, err := cmd.Output()

		if err != nil {
			return nil, err
		}

		exposedPort := strings.TrimSpace(strings.Split(string(output), ":")[1])
		intExposedPort, err := strconv.Atoi(exposedPort)

		if err != nil {
			return nil, err
		}

		ports = append(ports, PortMapping{intExposedPort, containerPort})
	}

	if err != nil {
		return nil, err
	}

	return ports, nil
}

func (tc *TestContainer) Stop() {
	log.Println("Stopping container")

	removeCommand := exec.Command("docker", "remove", "-f", tc.Name)

	err := removeCommand.Run()

	if err != nil {
		panic(err)
	}
}
