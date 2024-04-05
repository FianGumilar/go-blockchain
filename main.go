package main

import (
	"encoding/json"
	"go-blockchain/domain"
	"log"
)

func main() {
	bc := domain.NewBlockchain()

	bc.GiveMandate("Owner", "Investor", 1)

	if !bc.GiveMandate("Investor", "Quena", 1) {
		log.Println("insufficient-mandate")
	}
	bc.CreateBlock(bc.LatestBlock().Hash())

	if !bc.GiveMandate("Investor", "Ansuk", 1) {
		log.Println("insufficient-mandate")
	}
	bc.CreateBlock(bc.LatestBlock().Hash())

	log.Printf("Quena => :%v", bc.CalculateMandate("Quena"))
	log.Printf("Ansuk => :%v", bc.CalculateMandate("Ansuk"))

	data, _ := json.Marshal(bc)
	log.Println(string(data))
}
