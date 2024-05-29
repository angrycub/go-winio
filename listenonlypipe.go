//go:build windows
// +build windows

package winio

import (
	"net"
)

// ListenOnlyPipe creates a listener on a Windows named pipe path, e.g. \\.\pipe\mypipe.
// The pipe must already exist.
func ListenOnlyPipe(path string, c *PipeConfig) (net.Listener, error) {
	if c == nil {
		c = &PipeConfig{}
	}
	h, err := makeServerPipeHandle(path, nil, c, false)
	if err != nil {
		return nil, err
	}
	l := &win32PipeListener{
		firstHandle: h,
		path:        path,
		config:      *c,
		acceptCh:    make(chan (chan acceptResponse)),
		closeCh:     make(chan int),
		doneCh:      make(chan int),
	}
	go l.listenerRoutine()
	return l, nil
}
