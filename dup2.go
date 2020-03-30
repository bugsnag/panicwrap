// +build !arm64

package panicwrap

import (
	"golang.org/x/sys/unix"
)

func dup2(oldfd, newfd int) error {
	return unix.Dup2(oldfd, newfd)
}
