package cryptohedge

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/cryptohazard/coinmarketcap"
)

func ParseJSON(portfolioFile string, hedgeFile string) (*Cryptofolio, *Cryptohedge) {
	var crypto = new(Cryptofolio)
	var hedge = new(Cryptohedge)

	rawJSON, err := ioutil.ReadFile(portfolioFile)
	if err != nil {
		log.Fatal("Error in portfolio File ", portfolioFile, ": ", err)
	}
	json.Unmarshal(rawJSON, &crypto.CryptoArray)

	rawJSON, err = ioutil.ReadFile(hedgeFile)
	if err != nil {
		log.Fatal("Error in hedge fund File ", hedgeFile, ": ", err)
	}

	json.Unmarshal(rawJSON, &hedge.ShareArray)

	return crypto, hedge

}

func GetRate(crypto *Cryptofolio) error {

	s := make([]string, len(crypto.CryptoArray))
	for _, c := range crypto.CryptoArray {
		s = append(s, c.Name)
	}

	ticker, err := coinmarketcap.GetData(s)
	//fmt.Println(ticker)
	if err != nil {
		fmt.Println("error ticker")
		return err
	}
	for _, c := range crypto.CryptoArray {
		c.Rate = ticker.Coins[c.Name].PriceEUR
	}

	return nil
}
