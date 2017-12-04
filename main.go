package main

// not very important NOTES!:
// byte is an alias of uint8.
// uint8 is an unsigned integer type that is at leat 8 bits in size. range[0-255].
// [] this is a Slice in GoLang.
// int64 is the set of all signed 64-bit integers. range[-9223372036854775808,9223372036854775807].

import (
	"fmt" // Implements formatted I/O with functions analogus to C's printf and scanf.
	"strconv"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Alex")
	bc.AddBlock("Send 2 more BTC to Alex")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x \n", block.PrevBlockHash) // %x base 16, with lower-case letters for a-f
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
