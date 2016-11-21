package main

import (
	"flag"
	"github.com/rmrobinson/bottlerocket-go"
	"fmt"
)

var (
	devicePath = flag.String("devicePath", "/dev/firecracker", "The path to the serial device to communicate over")
	address = flag.String("address", "F1", "The address of the device to control")
	command = flag.String("command", "", "The command to send (ON or OFF currently supported")
)

func main() {
	flag.Parse()

	var br bottlerocket_go.Bottlerocket

	err := br.Open(*devicePath)

	if err != nil {
		fmt.Printf("Unable to initialize bottlerocket: %s\n", err.Error())
		return
	}

	defer br.Close()

	err = br.SendCommand(*address, *command)

	if err != nil {
		fmt.Printf("Error sending command %s to device %s: %s\n", *command, *address, err.Error())
	} else {
		fmt.Printf("%s set to %s\n", *address, *command)
	}
}
