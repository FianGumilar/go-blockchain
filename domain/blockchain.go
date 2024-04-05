package domain

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Blockchain struct {
	Pool  []*Mandate `json:"pool"`
	Chain []*Block   `json:"chain"`
}

func NewBlockchain() *Blockchain {
	bc := LoadDatabase()
	if len(bc.Chain) == 0 {
		bc.CreateGenesis()
		bc.CreateBlock(fmt.Sprintf("%x", [32]byte{}))
	}

	return &bc
}

func LoadDatabase() Blockchain {
	f, err := os.OpenFile("database/blockchain.db", os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	blockchain := Blockchain{}

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			os.Exit(1)
		}

		var blockSerialized BlockSerialized
		err = json.Unmarshal(scanner.Bytes(), &blockSerialized)
		if err != nil {
			os.Exit(1)
		}

		if len(blockchain.Chain) > 0 && (blockchain.LatestBlock().Hash() != blockSerialized.Value.Header.PrevHash) {
			log.Fatal("invalid block-chain database")
		}

		blockchain.Chain = append(blockchain.Chain, blockSerialized.Value)
	}
	return blockchain
}

func (bc *Blockchain) CreateBlock(prevHash string) *Block {
	b := NewBlock(prevHash, bc.Pool)
	bc.Chain = append(bc.Chain, b)
	bc.Pool = []*Mandate{}

	return b
}

func (bc *Blockchain) GiveMandate(from, to string, value int64) bool {
	if bc.CalculateMandate(from) < value {
		return false
	}

	m := NewMandate(from, to, value)
	bc.Pool = append(bc.Pool, m)
	return true
}

func (bc *Blockchain) LatestBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) CreateGenesis() {
	m := NewMandate("Owner", "Investor", 10)
	bc.Pool = append(bc.Pool, m)
}

func (bc *Blockchain) CalculateMandate(user string) int64 {
	var total int64

	for _, v := range bc.Chain {
		for _, v2 := range v.Mandates {
			if v2.To == user {
				total += v2.Value
			}

			if v2.From == user {
				total -= v2.Value
			}
		}
	}
	return total
}
