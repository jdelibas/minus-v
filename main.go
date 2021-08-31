package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func checkBinaryVersion(app string, arg string) (version string, err error) {
	cmd := exec.Command(app, arg)
	stdout, err := cmd.Output()

	if err != nil {
		return "", errors.New("Nope" + arg + "didnt work")
	}
	return string(stdout), nil
}

func checkPacmanVersion(app string) (version string, err error) {
	cmd := exec.Command("pacman -Qi", app)
	stdout, err := cmd.Output()

	if err != nil {
		return "", errors.New("Nope pacman didnt work")
	}
	return string(stdout), nil
}

func main() {
	app := os.Args[1]

	foundVersion := false
	args := []string{"-v", "--version", "-V", "version"}

	for _, arg := range args {
		res, err := checkBinaryVersion(app, arg)
		if err != nil {
			continue
		}
		fmt.Printf(res)
		foundVersion = true
		break
	}

	res, err := checkPacmanVersion(app)
	if err != nil {
		fmt.Println(res)
		return
	}

	if foundVersion != true {
		fmt.Println("Could not get version for " + app)
	}
}
