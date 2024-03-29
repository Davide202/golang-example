package main

import (
	"log"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	routineGroup()
}

func routineGroup() {
	log.Println("Start...")
	defer log.Println("...Stop")
	wg.Add(3) //aggiungo 3 routine
	go count(1)
	go count(2)
	go count(3)
	wg.Wait() //aspetto che tutte le routine finiscano
}

func count(seconds int) {
	d := time.Duration(seconds) * time.Second
	for i := 0; i < 3; i++ {
		log.Println("                    ROUTINE [ " + d.String() + " ] STEP: " + strconv.Itoa(i+1))
		time.Sleep(d)
	}
	wg.Done() //comunico al gruppo che questa routine ha terminato
}
