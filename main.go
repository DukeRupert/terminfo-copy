package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Create a reader for user input
	reader := bufio.NewReader(os.Stdin)

	// Prompt for remote server address
	fmt.Print("Enter remote server address (e.g., user@hostname): ")
	server, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Clean the server input by removing whitespace
	server = strings.TrimSpace(server)

	// Validate server input
	if server == "" {
		fmt.Fprintf(os.Stderr, "Server address cannot be empty\n")
		os.Exit(1)
	}

	// Execute infocmp command and get its output
	infocmpCmd := exec.Command("infocmp", "-x")
	infocmpOutput, err := infocmpCmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing infocmp: %v\n", err)
		os.Exit(1)
	}

	// Prepare ssh command
	sshCmd := exec.Command("ssh", server, "tic", "-x", "-")
	
	// Set up pipe to send infocmp output to ssh
	sshCmd.Stdin = strings.NewReader(string(infocmpOutput))
	sshCmd.Stdout = os.Stdout
	sshCmd.Stderr = os.Stderr

	// Execute the ssh command
	err = sshCmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing ssh command: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully copied Terminfo to %s\n", server)
}
