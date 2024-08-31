package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

const serverAddr = "127.0.0.1:8080"

func main() {
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		panic(err)
	}

	var s http.Server

	mux := new(http.ServeMux)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "[server1] Hello, world!")
	})
	mux.HandleFunc("/passfd", func(w http.ResponseWriter, r *http.Request) {
		if err := sendListener(lis.(*net.TCPListener)); err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}

		fmt.Fprintf(w, "Success")
		time.AfterFunc(time.Millisecond*50, func() {
			log.Println("Shutdown server ...")
			s.Shutdown(context.Background())
		})
	})

	s.Handler = mux
	log.Printf("Server is listening on %s ...\n", serverAddr)
	s.Serve(lis)
	log.Println("Bye bye")
}

func sendListener(lis *net.TCPListener) error {
	// connect to the unix socket
	const udsPath = "/tmp/fd-pass-example.sock"
	conn, err := net.Dial("unix", udsPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	connFd, err := getConnFd(conn.(*net.UnixConn))
	if err != nil {
		return err
	}

	// pass listener fd
	lisFd, err := getConnFd(lis)
	if err != nil {
		return err
	}
	rights := unix.UnixRights(int(lisFd))
	return unix.Sendmsg(connFd, nil, rights, nil, 0)
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
