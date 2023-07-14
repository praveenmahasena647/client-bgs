package helpers

import (
	"io"
	"net"
	"os"
	"os/exec"
)

func DialClient() error {
	var TCPClient, dialErr = net.Dial("tcp", ":42069")

	if dialErr != nil {
		return dialErr
	}

	defer TCPClient.Close()

	var fileBytes, fileErr = io.ReadAll(TCPClient)

	if fileErr != nil {
		return fileErr
	}

	var bgResult = changeBg(fileBytes)

	if bgResult != nil {
		return bgResult
	}

	return nil
}

func changeBg(b []byte) error {
	var file, fileErr = os.Create("bg.png")

	if fileErr != nil {
		return fileErr
	}

	var _, writeErr = file.Write(b)

	if writeErr != nil {
		return writeErr
	}

	defer file.Close()

	var cmd = exec.Command("feh", "./bg.png", "-F", "--bg-max", "--bg-c")

	var _, cmdErr = cmd.Output()

	if cmdErr != nil {
		return cmdErr
	}
	return nil
}
