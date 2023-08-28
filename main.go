package main

import (
	"bufio"
	"go-sdr/util"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func onMessage(message string) {
	println("raw message: ", message)
	if !strings.Contains(message, "ALN") {
		return
	}
	newMessage := util.MessageFromString(message)
	println("new message: ", util.MessageDebugString(newMessage))
}

func main() {
	command := "rtl_fm -f 169.65M -M fm -s 22050 -p 45 -g 30 | multimon-ng -a FLEX -t raw -"
	cmd := exec.Command("sh", "-c", command)
	cmd.Stderr = os.Stderr
	sdrPipe, _ := cmd.StdoutPipe()

	caplist := make(map[string]string)
	capcodesFile, _ := os.Open("capcodes.dict")
	defer capcodesFile.Close()
	scanner := bufio.NewScanner(capcodesFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		if len(parts) == 2 {
			caplist[parts[0]] = parts[1]
		}
	}

	_ = cmd.Start()

	scanner = bufio.NewScanner(sdrPipe)
	for scanner.Scan() {
		p2000 := scanner.Text()

		onMessage(p2000)
	}

	cmd.Process.Signal(syscall.SIGKILL)
	cmd.Wait()
}
