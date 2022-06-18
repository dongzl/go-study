package signal

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"testing"
)

// kill -2 进程号 是 Interrupt
// kill -15 进程号 是SIGTERM
// signal.Notify(c, os.Interrupt, syscall.SIGTERM) 注册2个信号
func TestName(t *testing.T) {
	pid := os.Getpid()
	fmt.Printf("进程 PID: %d \n", pid)
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		aa := <-c
		fmt.Println(aa)
		fmt.Println("catch signal")
		cancel()
		//<-c
		//os.Exit(1) // second signal. Exit directly.
	}()
	select {
	case <-ctx.Done():
		fmt.Println("111")
		return
	}
}

func TestPid(t *testing.T) {
	pid := os.Getpid()
	fmt.Printf("进程 PID: %d \n", pid)

	prc := exec.Command("ps", "-p", strconv.Itoa(pid), "-v")
	out, err := prc.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
