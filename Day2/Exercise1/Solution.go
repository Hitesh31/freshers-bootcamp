package main

import "fmt"

func Frequency(s string) map[rune]int {
	m := map[rune]int{}
	for _, r := range s {
		m[r]++
	}
	return m
}
func countFreq(stream []string) map[rune]int {

	sum := map[rune]int{}
	count := len(stream)
	results := make(chan map[rune]int, count)

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

	return sum
}
func main() {
	stream := make([]string, 5)
	stream[0]="quick"
	stream[1]="brown"
	stream[2]="fox"
	stream[3]="lazy"
	stream[4]="dog"
	fmt.Println(countFreq(stream))
}
