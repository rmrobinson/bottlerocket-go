package bottlerocket_go

/*
#cgo LDFLAGS: -lbr
#include <fcntl.h>
#include <errno.h>
#include <unistd.h>
#include <br_cmd.h>

int open_gowrapper(const char* path) {
	return open(path, O_RDONLY | O_NONBLOCK);
}
*/
import "C"
import (
	"errors"
	"strconv"
)

type Bottlerocket struct {
	fd   C.int
	path string
}

func (br *Bottlerocket) Open(path string) error {
	if len(br.path) > 0 {
		return errors.New("Bottlerocket already set up")
	}

	ret, err := C.open_gowrapper(C.CString(path))

	if ret < 0 {
		return err
	}

	br.fd = ret
	br.path = path

	return nil
}

func (br *Bottlerocket) Close() {
	if len(br.path) > 0 {
		C.close(br.fd)
		br.path = ""
	}
}

func (br *Bottlerocket) Path() string {
	return br.path
}

func (br *Bottlerocket) SendCommand(address string, command string) error {
	if len(br.path) < 1 {
		return errors.New("Bottlerocket not set up")
	}

	var cmd C.int

	// TODO: support more commands
	switch command {
	case "ON":
		cmd = C.ON
	case "OFF":
		cmd = C.OFF
	default:
		return errors.New("Invalid command specified")
	}

	var addr C.uchar

	if len(address) != 2 && len(address) != 3 {
		return errors.New("Invalid address specified (to many or few parts)")
	}

	if address[0] < 'A' || address[0] > 'P' {
		return errors.New("Invalid address specified (house invalid)")
	}

	house := int(address[0] - 'A')
	device, err := strconv.Atoi(address[1:])

	if err != nil {
		return err
	}

	if device < 1 || device > 16 {
		return errors.New("Invalid adress specified (device invalid)")
	}

	addr = C.uchar(house<<4 | device - 1)

	_, err = C.br_cmd(br.fd, addr, cmd)

	return err
}

