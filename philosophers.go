package main

import (
	"fmt"
	"slices"
	"sync"
	"time"
)

var wg sync.WaitGroup
var wgHost sync.WaitGroup

type ChopStick struct {
	mut sync.Mutex
}

type Philosopher struct {
	RightStick *ChopStick
	LeftStick  *ChopStick
	GetPerm    chan bool
	ID         int
}

func (philosopher *Philosopher) Eat(perm chan int, wait chan bool, activePhilos *[]int) {
	defer close(philosopher.GetPerm)
	defer wg.Done()

	for i := 0; i < 3; i++ {
		// request permission
		perm <- philosopher.ID

		// wait for permission to eat
		<-philosopher.GetPerm

		*activePhilos = append(*activePhilos, philosopher.ID)

		philosopher.LeftStick.mut.Lock()
		philosopher.RightStick.mut.Lock()
		fmt.Printf("Starting to eat %v \n", philosopher.ID)
		time.Sleep(1 * time.Second)
		fmt.Printf("Finishing to eat %v \n", philosopher.ID)
		philosopher.RightStick.mut.Unlock()
		philosopher.LeftStick.mut.Unlock()

		(*activePhilos) = slices.DeleteFunc(*activePhilos, func(e int) bool { return e == philosopher.ID })

		<-wait
	}
}

type Host struct {
	Permission   chan int
	Philosophers []*Philosopher
	Wait         chan bool
	ActivePhilos []int
}

func (host *Host) Manage(abort chan bool) {
	for {
		select {
		case <-abort:
			wgHost.Done()
			return
		case p := <-host.Permission:
			host.Wait <- true

			if len(host.ActivePhilos) == 0 {
				host.Philosophers[p].GetPerm <- true
			} else {
				currPhilo := host.ActivePhilos[0]
				for len(host.ActivePhilos) != 0 && (p == (currPhilo+1)%5 || p == (currPhilo-1)%5) {
					host.Permission <- p
					p = <-host.Permission
				}

				host.Philosophers[p].GetPerm <- true
			}
		}
	}
}

func main() {
	numPhilo := 5
	maxAllowedPhilo := 2

	sticks := make([]*ChopStick, numPhilo)
	philosophers := make([]*Philosopher, numPhilo)

	for i := 0; i < numPhilo; i++ {
		sticks[i] = new(ChopStick)
	}

	for i := 0; i < numPhilo; i++ {
		philosophers[i] = &Philosopher{
			ID:         i,
			LeftStick:  sticks[i],
			RightStick: sticks[(i+1)%numPhilo],
			GetPerm:    make(chan bool),
		}
	}

	host := &Host{
		Permission:   make(chan int, numPhilo+1),
		Philosophers: philosophers,
		ActivePhilos: make([]int, 0, maxAllowedPhilo),
		Wait:         make(chan bool, maxAllowedPhilo),
	}

	abort := make(chan bool, 1)

	wgHost.Add(1)

	go host.Manage(abort)

	for i := 0; i < numPhilo; i++ {
		wg.Add(1)

		go philosophers[i].Eat(host.Permission, host.Wait, &host.ActivePhilos)
	}

	wg.Wait()
	abort <- true

	wgHost.Wait()
}
