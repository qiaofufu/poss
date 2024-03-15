package main

import (
	"github.com/klauspost/reedsolomon"
	"os"
)

func main() {
	encoder, err := reedsolomon.New(10, 3)
	if err != nil {
		panic(err)
	}
	// 分配输入大小
	data := encoder.(reedsolomon.Extensions).AllocAligned(5000)
	open, err := os.Open("1")
	stat, err := os.Stat("1")
}
