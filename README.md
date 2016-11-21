# bottlerocket-go

A golang library which wraps the Linuxha Bottlerocket C library.

The library currently works with the 0.05b3 build available at https://github.com/linuxha/bottlerocket, and assumes that the lib_install target has been run.

## Usage

Create an instance of the bottlerocket-go.Bottlerocket struct, and then call Open with the path of the serial port that the Firecracker module is plugged in to.
Make sure to defer calling Close() so things are cleaned up properly.
Commands are sent by calling SendCommand with the desired address and command. At the moment only ON and OFF are supported as commands.

## TODO

* godoc existing code

* Support a larger number of commands