package main

import (
	"fmt"
	"time"

	"github.com/carvalhocaio/go-search-engine/internal/fetcher"
	"github.com/carvalhocaio/go-search-engine/internal/models"
	"github.com/carvalhocaio/go-search-engine/internal/processor"
)

func main() {
	start := time.Now()

	priceChannel := make(chan models.PriceDetail)
	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAVG(priceChannel, done)

	<-done

	fmt.Printf("\nTempo total: %s\n", time.Since(start))
}
