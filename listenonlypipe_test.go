//go:build windows
// +build windows

package winio

import (
	"bufio"
	"testing"
)

func TestMultiListener(t *testing.T) {
	l1, err := ListenPipe(testPipeName, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer l1.Close()

	l2, err := ListenOnlyPipe(testPipeName, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer l2.Close()

	ch1 := make(chan int)
	go server(l1, ch1)

	ch2 := make(chan int)
	go server(l2, ch2)

	c, err := DialPipe(testPipeName, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c))
	_, err = rw.WriteString("hello world\n")
	if err != nil {
		t.Fatal(err)
	}
	err = rw.Flush()
	if err != nil {
		t.Fatal(err)
	}

	s, err := rw.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	ms := "got hello world\n"
	if s != ms {
		t.Errorf("expected '%s', got '%s'", ms, s)
	}

	<-ch1
	<-ch2
}
