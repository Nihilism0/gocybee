//实在搞不懂怎么输入q停止打炮
package main

import (
	"fmt"
	"sync"
	"time"
)

var relx1 sync.RWMutex
var relx2 sync.RWMutex
var relx3 sync.RWMutex

func ready(ch1 chan bool) {
	for {
		time.Sleep(5 * time.Second)
		relx1.Lock()
		if len(ch1) == 0 {
			ch1 <- true
			{
				fmt.Printf("装填->")
			}
		}
		relx1.Unlock()
	}
}

func see(ch1 chan bool, ch2 chan bool) {
	for {
		time.Sleep(6 * time.Second)
		relx2.Lock()
		if len(ch1) == 1 && len(ch2) == 0 {
			ch2 <- true
			fmt.Printf("瞄准->")
		}
		relx2.Unlock()
	}
}
func shoot(ch1 chan bool, ch2 chan bool) {
	for {
		time.Sleep(7 * time.Second)
		relx3.Lock()
		if len(ch1) == 1 && len(ch2) == 1 {
			fmt.Printf("发射!\n")
			<-ch1
			<-ch2
		}
		relx3.Unlock()
	}
}
func main() {

	ch1 := make(chan bool, 1)
	ch2 := make(chan bool, 1)
	go ready(ch1)
	go ready(ch1)
	go ready(ch1)
	go ready(ch1)
	go ready(ch1)
	go ready(ch1)
	go ready(ch1)
	go ready(ch1)
	go ready(ch1)
	go ready(ch1)
	go see(ch1, ch2)
	go see(ch1, ch2)
	go see(ch1, ch2)
	go see(ch1, ch2)
	go see(ch1, ch2)
	go shoot(ch1, ch2)
	go shoot(ch1, ch2)
	go shoot(ch1, ch2)
}
