package main

import (
	"bufio"
	"fmt"
	"github.com/bsphere/le_go"
	"os"
)

func main() {
	le_token := os.Getenv("LE_TOKEN")
	if le_token == "" {
		exitWithStdErr("Please export LE_TOKEN")
	}

	le, err := le_go.Connect(le_token)
	if err != nil {
		exitWithStdErr("Failed to connect Logentries")
	}

	defer le.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		le.Println(scanner.Text())
	}
}

func exitWithStdErr(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
