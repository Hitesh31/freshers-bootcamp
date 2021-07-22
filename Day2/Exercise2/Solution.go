package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func expectedRating(val float64, wg *sync.WaitGroup, sum *float64){
	defer wg.Done()
	time.Sleep(time.Second*time.Duration(rand.Intn(2)))
	*sum=*sum+val
}
func main () {

	var wg sync.WaitGroup
	var students[200] float64
	k:=0.0
	for i:=0;i<200;i++ {
		students[i]=float64(rand.Intn(5))
		k=k+students[i]
	}
	sum := 0.0
	for i:=0;i<200;i++ {
		wg.Add(1)
		go expectedRating(students[i], &wg, &sum)
	}
	wg.Wait()
	sum=sum/200
	fmt.Println(sum)
	fmt.Println(k)
}
