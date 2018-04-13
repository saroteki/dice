package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := getRand()

	luck := shakeDice(r)
	fmt.Printf(luck)
}

func getRand() int {
	t := time.Now().UnixNano()
	rand.Seed(t)
	return rand.Intn(6) + 1
}

func shakeDice(num int) (luck string) {
	switch num {
	case 6:
		return "大吉"
	case 5, 4:
		return "中吉"
	case 3, 2:
		return "吉"
	case 1:
		return "凶"
	default:
		return "Unknown Number"
	}
}
