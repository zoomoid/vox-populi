package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	c1 := make(chan error, 1)
	go faultyBackend(c1)

	select {
	case err := <-c1:
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Print("success")
		}
	case <-time.After(2 * time.Second):
		fmt.Print("timeout")
	}
}

func healthyBackend(c chan error) {
	time.Sleep(200 * time.Millisecond)
	c <- nil
}

func slowBackend(c chan error) {
	time.Sleep(3 * time.Second)
	c <- nil
}

func slowFaultyBackend(c chan error) {
	time.Sleep(3 * time.Second)
	c <- errors.New("error after time")
}

func faultyBackend(c chan error) {
	time.Sleep(1 * time.Second)
	c <- errors.New("immediate error")
}
