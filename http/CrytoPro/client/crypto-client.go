package client

import (
	"encoding/json"
	"log"
	"net/http"

	"example.local/demo/Desktop/goProgramming/goLanguage/http/CrytoPro/model"
)

func FetchCrypto(fiat string, crypto string) (string, error) {
	url := "https://api.nomics.com/v1/currencies/ticker?key=3990ec554a414b59dd85d29b2286dd85&interval=1d&ids=" + crypto + "&convert=" + fiat
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("ooops an error occured please try again")
	}
	defer resp.Body.Close()
	var cResp model.Cryptoresponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Println("Oops! an error occured ", err)
	}
	return cResp.TextOutput(), nil
}
