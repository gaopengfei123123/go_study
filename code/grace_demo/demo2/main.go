package main

// 代码来源 https://studygolang.com/articles/14038?fr=sidebar
import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

type listener struct {
	Addr     string `json:"addr"`
	FD       int    `json:"fd"`
	Filename string `json:"filename"`
}

func importListener(addr string) (net.Listener, error) {
	// 从环境变量中抽离出被编码的 listener 的元数据。
	listenerEnv := os.Getenv("LISTENER")
	if listenerEnv == "" {
		return nil, fmt.Errorf("unable to find LISTENER environment variable")
	}

	// 解码 listener 的元数据。
	var l listener
	err := json.Unmarshal([]byte(listenerEnv), &l)
	if err != nil {
		return nil, err
	}
	if l.Addr != addr {
		return nil, fmt.Errorf("unable to find listener for %v", addr)
	}

	// 文件已经被传入到这个进程中，从元数据中抽离文件描述符和名字，为 listener 重建/发现 *os.file
	listenerFile := os.NewFile(uintptr(l.FD), l.Filename)
	if listenerFile == nil {
		return nil, fmt.Errorf("unable to create listener file: %v", err)
	}
	defer listenerFile.Close()

	// Create a net.Listener from the *os.File.
	ln, err := net.FileListener(listenerFile)
	if err != nil {
		return nil, err
	}

	return ln, nil
}

func createListener(addr string) (net.Listener, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return ln, nil
}

func createOrImportListener(addr string) (net.Listener, error) {
	// 尝试为地址导入一个 listener, 如果导入成功，则使用。
	ln, err := importListener(addr)
	if err == nil {
		fmt.Printf("Imported listener file descriptor for %v.\n", addr)
		return ln, nil
	}
	// 没有 listener 被导入，这就意味着进程必须自己创建一个。
	ln, err = createListener(addr)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Created listener file descriptor for %v.\n", addr)
	return ln, nil
}

func getListenerFile(ln net.Listener) (*os.File, error) {
	switch t := ln.(type) {
	case *net.TCPListener:
		return t.File()
	case *net.UnixListener:
		return t.File()
	}
	return nil, fmt.Errorf("unsupported listener: %T", ln)
}

func forkChild(addr string, ln net.Listener) (*os.Process, error) {
	// 从 listener 中获取文件描述符，在环境变量编码在传递给这个子进程的元数据。
	lnFile, err := getListenerFile(ln)
	if err != nil {
		return nil, err
	}
	defer lnFile.Close()
	l := listener{
		Addr:     addr,
		FD:       3,
		Filename: lnFile.Name(),
	}
	listenerEnv, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}
	// 将 stdin, stdout, stderr 和 listener 传入子进程。
	// 译注: 以上四个文件描述符分别为 0,1,2,3
	files := []*os.File{
		os.Stdin,
		os.Stdout,
		os.Stderr,
		lnFile,
	}

	// 获取当前环境变量，并且传入子进程。
	environment := append(os.Environ(), "LISTENER="+string(listenerEnv))

	// 获取当前进程名和工作目录
	execName, err := os.Executable()
	if err != nil {
		return nil, err
	}
	execDir := filepath.Dir(execName)

	// 生成子进程
	p, err := os.StartProcess(execName, []string{execName}, &os.ProcAttr{
		Dir:   execDir,
		Env:   environment,
		Files: files,
		Sys:   &syscall.SysProcAttr{},
	})
	if err != nil {
		return nil, err
	}

	return p, nil
}

func waitForSignals(addr string, ln net.Listener, server *http.Server) error {
	signalCh := make(chan os.Signal, 1024)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGUSR2, syscall.SIGINT, syscall.SIGQUIT)
	for {
		select {
		case s := <-signalCh:
			fmt.Printf("%v signal received.\n", s)
			switch s {
			case syscall.SIGHUP:
				// Fork 一个子进程。
				p, err := forkChild(addr, ln)
				if err != nil {
					fmt.Printf("Unable to fork child: %v.\n", err)
					continue
				}
				fmt.Printf("Forked child %v.\n", p.Pid)

				// 创建一个在 5 秒钟过去的 Context, 使用这个超时定时器关闭。
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				// 返回关闭过程中发生的任何错误。
				return server.Shutdown(ctx)
			case syscall.SIGUSR2:
				// Fork 一个子进程。
				p, err := forkChild(addr, ln)
				if err != nil {
					fmt.Printf("Unable to fork child: %v.\n", err)
					continue
				}
				// 输出被 fork 的子进程的 PID，并等待更多的信号。
				fmt.Printf("Forked child %v.\n", p.Pid)
			case syscall.SIGINT, syscall.SIGQUIT:
				// 创建一个在 5 秒钟过去的 Context, 使用这个超时定时器关闭。
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				// 返回关闭过程中发生的任何错误。
				return server.Shutdown(ctx)
			}
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %v!\n", os.Getpid())
}

func startServer(addr string, ln net.Listener) *http.Server {
	http.HandleFunc("/hello", handler)

	httpServer := &http.Server{
		Addr: addr,
	}
	go httpServer.Serve(ln)

	return httpServer
}

func main() {
	// Parse command line flags for the address to listen on.
	var addr string
	flag.StringVar(&addr, "addr", ":8080", "Address to listen on.")

	// Create (or import) a net.Listener and start a goroutine that runs
	// a HTTP server on that net.Listener.
	ln, err := createOrImportListener(addr)
	if err != nil {
		fmt.Printf("Unable to create or import a listener: %v.\n", err)
		os.Exit(1)
	}
	server := startServer(addr, ln)

	// 等待复制或结束的信号
	err = waitForSignals(addr, ln, server)
	if err != nil {
		fmt.Printf("Exiting: %v\n", err)
		return
	}
	fmt.Printf("Exiting.\n")
}
