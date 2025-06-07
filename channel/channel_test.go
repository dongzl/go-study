package channel

import (
	"fmt"
	"testing"
	"time"
)

func Test_channel_1(t *testing.T) {
	st := time.Now()
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Second * 2)
		<-ch
	}()
	ch <- true
	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds())
	time.Sleep(time.Second * 5)
}

func Test_channel_2(t *testing.T) {
	st := time.Now()
	ch := make(chan bool, 2)
	go func() {
		time.Sleep(time.Second * 2)
		<-ch
	}()
	ch <- true
	ch <- true
	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds())
	ch <- true
	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds())
	time.Sleep(time.Second * 5)
}
