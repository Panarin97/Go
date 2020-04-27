package main

import (
	"fmt"
	"sync"
	"math/rand"
)

type ChopS struct {
	sync.Mutex
}

type Phil struct {
	LeftChopS, RightChopS *ChopS
	number int
	host chan bool
}

func (p Phil) eat(wg *sync.WaitGroup) {
	defer wg.Done()

	// randomizing the process
	r := rand.Intn(2)
	r_1 := rand.Intn(2)

	for i := 1; i <= 3; i++ {

		p.host <- true

		// random pick up 
		if (r == 0) {
			p.LeftChopS.Lock()
			p.RightChopS.Lock()

			fmt.Println("Starting to eat", p.number)

		} else {
			p.RightChopS.Lock()
			p.LeftChopS.Lock()

			fmt.Println("Starting to eat", p.number)
		}

		// random put down
		if (r_1 == 0) {
			fmt.Println("Finished eating", p.number)
		
			p.LeftChopS.Unlock()
			p.RightChopS.Unlock()

		} else {
			fmt.Println("Finished eating", p.number)

			p.RightChopS.Unlock()
			p.LeftChopS.Unlock()
		}

		<- p.host
	}
}


func main() {

	host_channel := make(chan bool, 2)

	// creating 5 chopsticks
	chops := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		chops[i] = new(ChopS)
	}

	// creating 5 philosophers
	phils := make([]*Phil, 5)
	for i := 0; i < 5; i++ {
		phils[i] = &Phil{chops[i], chops[(i+1)%5], i+1, host_channel}

	}

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go phils[i].eat(&wg)
	} 

	wg.Wait()
}
