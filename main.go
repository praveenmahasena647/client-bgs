package main

import (
	"io/ioutil"
	"net"
	"os"
	"os/exec"
)

func main() {
	var con, conErr = net.Dial("tcp", ":42069")
	if conErr != nil {
		panic("couldnt do the connections")
	}

	defer con.Close()

	var data, dataErr = ioutil.ReadAll(con)
	con.Write(data)
	if dataErr != nil {
		panic("data ioutil err")
	}

	var err error = drawImg(data)

	if err != nil {
		panic(err)
	}
}

func drawImg(data []byte) error {
	var file, err = os.Create("bg.png")

	if err != nil {
		return err
	}

	_, err = file.Write(data)

	if err != nil {
		return err
	}

	defer file.Close()
	var cmd = exec.Command("feh", "./bg.png", "-F", "--bg-max", "--bg-c")

	_, err = cmd.Output()

	return err
}
