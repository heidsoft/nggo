package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"bytes"
	"syscall"
	"time"
)

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

func Echo() {
	// Create an *exec.Cmd
	cmd := exec.Command("echo", "Called from Go!")

	// Combine stdout and stderr
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output) // => go version go1.3 darwin/amd64

}

func GoVersion() {
	// Create an *exec.Cmd
	cmd := exec.Command("go", "version")

	// Stdout buffer
	cmdOutput := &bytes.Buffer{}
	// Attach buffer to command
	cmd.Stdout = cmdOutput

	// Execute command
	printCommand(cmd)
	err := cmd.Run() // will wait for command to return
	printError(err)
	// Only output the commands stdout
	printOutput(cmdOutput.Bytes()) // => go version go1.3 darwin/amd64
}

func Ls() {
	cmd := exec.Command("ls", "/User")
	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		printError(err)
		// Did the command fail because of an unsuccessful exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	} else {
		// Command was successful
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
	}

}

func Cat() {
	cmd := exec.Command("cat", "/dev/random")
	//cmd := exec.Command("fping","-g","10.0.1.191/24")
	randomBytes := &bytes.Buffer{}
	cmd.Stdout = randomBytes

	// Start command asynchronously
	err := cmd.Start()
	printError(err)

	// Create a ticker that outputs elapsed time
	ticker := time.NewTicker(time.Second)
	go func(ticker *time.Ticker) {
		now := time.Now()
		for _ = range ticker.C {
			printOutput(
				[]byte(fmt.Sprintf("%s", time.Since(now))),
			)
		}
	}(ticker)

	// Create a timer that will kill the process
	timer := time.NewTimer(time.Second * 4)
	go func(timer *time.Timer, ticker *time.Ticker, cmd *exec.Cmd) {
		for _ = range timer.C {
			err := cmd.Process.Signal(os.Kill)
			printError(err)
			ticker.Stop()
		}
	}(timer, ticker, cmd)

	// Only proceed once the process has finished
	cmd.Wait()
	printOutput(
		[]byte(fmt.Sprintf("%d bytes generated!", len(randomBytes.Bytes()))),
	)
}

func Fping() {
	cmd := exec.Command("fping", "10.0.1.191")
	// Combine stdout and stderr
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output)

}

func main() {
	Echo()
	GoVersion()
	Ls()
	Cat()
	Fping()
}