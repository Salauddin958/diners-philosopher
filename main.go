package main

import (
	"fmt"
	"sync"
	"time"
)

// five philisophers sitting around the table , only 2 can eat at a time
// since only 2 can eat at a time , using a buffered channel for blocking

var host = make(chan bool, 2)
var wg sync.WaitGroup

type ChopStick struct {
	sync.Mutex
}

type Philisopher struct {
	id        int
	leftChop  *ChopStick
	rightChop *ChopStick
}

func (p Philisopher) eat() {

	// for infinite times eating replace below line with "for true"
	for i := 0; i < 2; i++ {
		host <- true
		p.leftChop.Lock()
		p.rightChop.Lock()
		fmt.Println("eating started by ", p.id+1)
		time.Sleep(2 * time.Second)
		fmt.Println("eating ended by ", p.id+1)
		p.leftChop.Unlock()
		p.rightChop.Unlock()
		<-host
	}
	wg.Done()
}

func main() {
	chopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopStick)
	}
	philisophers := make([]*Philisopher, 5)
	for i := 0; i < 5; i++ {
		philisophers[i] = &Philisopher{i, chopSticks[i], chopSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philisophers[i].eat()
	}
	wg.Wait()
}
