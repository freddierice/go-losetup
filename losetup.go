package losetup

import (
	"io/ioutil"
	"regexp"
)

var valid, checked bool

var losetupRegex = regexp.MustCompile("loop-control")

// checkValid checks to see if loop-control is loaded as a kernel module
func checkValid() {
	checked = true
	buf, err := ioutil.ReadFile("/proc/misc")
	if err != nil {
		return
	}
	valid = losetupRegex.Match(buf)
}

// Valid checks to see if losetup is an availability
func Valid() bool {
	if !checked {
		checkValid()
	}
	return valid
}
