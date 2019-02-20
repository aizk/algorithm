package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	data := []int{1, 2, 3, 4, 5}
	shuffle(data)
	fmt.Printf("%+v", data)
}

func shuffle(x []int) {
	length := len(x)
	if length == 0 {
		return
	}
	rand.Seed(time.Now().Unix())

	i := length - 1
	for i > 0  {
		// rand
		j := rand.Intn(i + 1)
		// exchange
		x[i], x[j] = x[j], x[i]
		i--
	}
}