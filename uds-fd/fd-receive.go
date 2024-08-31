package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

const udsPath = "/tmp/fd-pass-example.sock"

func main() {
	os.Remove(udsPath) //nolint: errcheck
	lis, err := net.Listen("unix", udsPath)
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	log.Println("Wait receiving listener ...")
	conn, err := lis.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	httpLis := receiveListener(conn.(*net.UnixConn))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "[server2] Hello, world!")
	})
	log.Printf("Server is listening on %s ...\n", httpLis.Addr())
	http.Serve(httpLis, nil)
}

func receiveListener(conn *net.UnixConn) net.Listener {
	connFd, err := getConnFd(conn)
	if err != nil {
		panic(err)
	}

	// receive socket control message
	b := make([]byte, unix.CmsgSpace(4))
	_, _, _, _, err = unix.Recvmsg(connFd, nil, b, 0)
	if err != nil {
		panic(err)
	}

	// parse socket control message
	cmsgs, err := unix.ParseSocketControlMessage(b)
	if err != nil {
		panic(err)
	}
	fds, err := unix.ParseUnixRights(&cmsgs[0])
	if err != nil {
		panic(err)
	}
	fd := fds[0]
	log.Printf("Got socket fd %d\n", fd)

	// construct net listener
	f := os.NewFile(uintptr(fd), "listener")
	defer f.Close()

	l, err := net.FileListener(f)
	if err != nil {
		panic(err)
	}
	return l
}

func getConnFd(conn syscall.Conn) (connFd int, err error) {
	var rawConn syscall.RawConn
	rawConn, err = conn.SyscallConn()
	if err != nil {
		return
	}

	err = rawConn.Control(func(fd uintptr) {
		connFd = int(fd)
	})
	return
}
