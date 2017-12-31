package main

import (
	"fmt"
	"github.com/toorop/go-bittrex"
	"os"
)

func main() {
	API_KEY := os.Getenv("API_KEY")
	API_SECRET := os.Getenv("API_SECRET")

	// Bittrex client
	bittrex := bittrex.New(API_KEY, API_SECRET)

	// Get markets
	markets, err := bittrex.GetMarkets()
	fmt.Println(err, markets)
}
