package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func expectedRating(wg *sync.WaitGroup, sum *float32){
	defer wg.Done()
	time.Sleep(time.Second*time.Duration(rand.Intn(2)))
	*sum=*sum+float32(rand.Intn(5))
}
func main () {

	var wg sync.WaitGroup
	var cumulativeRating float32
	for i:=0;i<200;i++ {
		wg.Add(1)
		go expectedRating(&wg, &cumulativeRating)
	}
	wg.Wait()
	avRating := cumulativeRating/200
	fmt.Println(cumulativeRating,avRating)
}
