package main

import (
	"fmt"
	"os/exec"
)

func IsNat(ip string, DockerHostPort int) bool {
	fmt.Printf("Testing if docker %s:%d is reachable...", ip, DockerHostPort)
	cmdstring := fmt.Sprintf("nc -w 10 %s %d < /dev/null", ip, DockerHostPort)
	command := exec.Command("/bin/sh", "-c", cmdstring)
	command.Start()
	if err := command.Wait(); err == nil {
		fmt.Printf("execute \"%s\" success\n", cmdstring)
	} else {
		fmt.Printf("execute \"%s\" failure, %s\n", cmdstring, err)
	}
}

func main() {
	private_ip := "127.0.0.1"
	public_ip := "106.39.119.67"
	DockerHostPort := 10250

	test_nc(private_ip, DockerHostPort)
	test_nc(public_ip, DockerHostPort)

}
