package main

import (
	"encoding/json"
	"fmt"
)

var sum = map[string]int{}

func Frequency(s string) map[string]int {
	m := map[string]int{}
	for _, r := range s {
		m[string(r)]++
	}
	return m
}
func countFreq(stream []string) {
	count := len(stream)
	results := make(chan map[string]int, count)

	for _, s := range stream {
		go func(s string) {
			results <- Frequency(s)
		}(s)
	}

	for i := 0; i < count; i++ {
		for r, freq := range <-results {
			sum[r] += freq
		}
	}
}
func showFreq() {
	k, err := json.MarshalIndent(sum,""," ")
	if err != nil {
		fmt.Println("Error while parsing json")
		fmt.Println(err)
	}
	fmt.Printf(string(k))
}

func main() {
	stream := []string{"quick","brown","fox","lazy","dog"}
	countFreq(stream)
	showFreq()
}
