package losetup

import (
	"fmt"
	"os"
)

// Device represents a loop device /dev/loop#
type Device uint64

// open returns a file handle to /dev/loop# and returns an error if it cannot
// be opened.
func (device Device) open() (*os.File, error) {
	return os.Open(device.String())
}

// Path returns the path to the loopback device
func (device Device) Path() string {
	return fmt.Sprintf(DeviceFormatString, device)
}
