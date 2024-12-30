package main

import (
	"fmt"
	"time"

	"github.com/carvalhocaio/go-search-engine/internal/fetcher"
	"github.com/carvalhocaio/go-search-engine/internal/processor"
)

func main() {
	start := time.Now()

	priceChannel := make(chan float64)
	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAVG(priceChannel, done)

	<-done

	fmt.Printf("\nTempo total: %s\n", time.Since(start))
}
