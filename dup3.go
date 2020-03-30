//+build linux,arm64

package panicwrap

import (
	"golang.org/x/sys/unix"
)

func dup2(oldfd, newfd int) error {
	return unix.Dup3(oldfd, newfd, 0)
}
