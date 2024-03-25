package main

import (
	"encoding/json"
	"go-blockchain/domain"
	"log"
)

func main() {
	bc := domain.NewBlockchain()

	bc.GiveMandate("BBCA", "Quuena", 1)
	bc.CreateBlock(bc.LatestBlock().Hash())

	data, _ := json.Marshal(bc)
	log.Println(string(data))
}
