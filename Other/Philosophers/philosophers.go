package main

import (
	"fmt"
	"sync"
)

var host = make(chan int, 2)
var wg sync.WaitGroup

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
	number          int
}

func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		host <- 1
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.number)

		p.leftCS.Unlock()
		p.rightCS.Unlock()

		fmt.Printf("finishing eating %d\n", p.number)

		<-host
	}
	wg.Done()
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1}
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	wg.Wait()
}
