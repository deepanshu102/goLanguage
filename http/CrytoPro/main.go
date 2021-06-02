package main

import (
	"flag"
	"fmt"
	"log"

	"example.local/demo/Desktop/goProgramming/goLanguage/http/CrytoPro/client"
)

func main() {
	fiatCurrency := flag.String(
		"fiat", "INR", "The name of the fiat currency you would like to know the price of your crypto in ",
	)
	nameOfCrypto := flag.String(
		"crypto", "BTC", "Input the name of the CryptoCurrency you would like to know the price of",
	)
	flag.Parse()
	crypto, err := client.FetchCrypto(*fiatCurrency, *nameOfCrypto)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(crypto)
}
