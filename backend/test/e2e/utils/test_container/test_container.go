package test_container

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

type PortMapping struct {
	HostPort      string
	ContainerPort string
}

type Environment struct {
	Key   string
	Value string
}

type ContainerInfo struct {
	Image        string
	PortMappings []PortMapping
	Environment  []Environment
	WaitLog      string
	StartTimeout time.Duration
}

var Debug = false

type TestContainer struct {
	Name string
	Info ContainerInfo
}

func (tc *TestContainer) Start() error {
	args := Args{"run", "--name", tc.Name}

	args.AddEnv(tc.Info.Environment)
	args.AddPortMappings(tc.Info.PortMappings)
	args.AddImageName(tc.Info.Image)

	command := exec.Command("docker", args...)

	pipe, err := command.StdoutPipe()
	if err != nil {
		return err
	}

	pipeErr, err := command.StderrPipe()
	if err != nil {
		return err
	}

	err = command.Start()
	if err != nil {
		return err
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
		return fmt.Errorf("Timed out waiting for container to start")
	}

	log.Printf("Started container: %s", tc.Name)

	return nil
}

func (tc *TestContainer) Stop() {
	log.Println("Stopping container")

	removeCommand := exec.Command("docker", "remove", "-f", tc.Name)

	err := removeCommand.Run()

	if err != nil {
		panic(err)
	}
}
