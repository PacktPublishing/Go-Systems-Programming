package main

import (
	"fmt"
	"sync"
	"time"
)

var Password = secret{counter: 1, password: "myPassword"}

type secret struct {
	sync.RWMutex
	counter  int
	password string
}

func Change(c *secret, pass string) {
	c.Lock()
	fmt.Println("LChange")
	time.Sleep(20 * time.Second)
	c.counter = c.counter + 1
	c.password = pass
	c.Unlock()
}

func Show(c *secret) string {
	fmt.Println("LShow")
	time.Sleep(time.Second)
	c.RLock()
	defer c.RUnlock()
	return c.password
}

func Counts(c secret) int {
	c.RLock()
	defer c.RUnlock()
	return c.counter
}

func main() {

	fmt.Println("Pass:", Show(&Password))
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("Go Pass:", Show(&Password))
		}()
	}

	go func() {
		Change(&Password, "123456")
	}()

	fmt.Println("Pass:", Show(&Password))
	time.Sleep(time.Second)
	fmt.Println("Counter:", Counts(Password))
}
