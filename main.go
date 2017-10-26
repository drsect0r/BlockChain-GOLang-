package main

// not very important NOTES!:
// byte is an alias of uint8.
// uint8 is an unsigned integer type that is at leat 8 bits in size. range[0-255].
// [] this is a Slice in GoLang.
// int64 is the set of all signed 64-bit integers. range[-9223372036854775808,9223372036854775807].

import (
	"bytes"         // Implements functions for the manipulation of byte slices.
	"crypto/sha256" // Implements SHA256 algorithms as defined in FIPS 180-4.
	"fmt"           // Implements formatted I/O with functions analogus to C's printf and scanf.
	"strconv"       // Implements conversion to and from string representations of basic data types.
	"time"          // Provides functionality for measuring and displaying time.
)

type Block struct { // BLOCK in blockchain it's block that store valuable information.
	Timestamp     int64  // Current Timestamp (when the block is created).
	Data          []byte // Is the actual valuable information containing in the block.
	PrevBlockHash []byte // Stores the hash of the previous block.
	Hash          []byte // Is the hash of the block.
}

// Take the block fields,
// concatenate them and calculate a
// SHA-256 hash on the concatenated combination.
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]

}

// Following a Golang convention,
// we'll implement a function that'll simplify the creation of a block.
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(), // Unix returns Unix time, the number of seconds elapsed since January 1, 1970 UTC. [int64]
		[]byte(data),
		prevBlockHash,
		[]byte{},
	}
	block.SetHash()
	return block
}

//In it's essence blockchain is just a database with certain structure.
// It's an ordered, back-linked list.
// Which means that blocks are stored in the insertion and
// that each block is linked to the previos one.
type Blockchain struct {
	blocks []*Block
}

//Let's make it possible to add blocks to it
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NOTES!:
// To add a new block we need an existing block.
// SO..... the first block is called "GENESIS BLOCK"
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// Now, we can implement a function that creates a
// blockchain with the genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Alex")
	bc.AddBlock("Send 2 more BTC to Alex")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x \n", block.PrevBlockHash) // %x base 16, with lower-case letters for a-f
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println(block.Timestamp)
		fmt.Println()
	}
}
