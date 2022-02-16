package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {

	for i := 0; i < 100; i++ {
		go func() {
			r, _ := rand.Int(rand.Reader, big.NewInt(5))
			fmt.Println(r.Int64() / 10)
		}()
	}

	time.Sleep(2 * time.Second)
}
