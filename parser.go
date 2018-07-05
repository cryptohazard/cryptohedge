package cryptohedge

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/cryptohazard/coinmarketcap"
)

func ParseJSON(portfolioFile string) *Cryptofolio {
	var crypto = new(Cryptofolio)
	rawJSON, err := ioutil.ReadFile(portfolioFile)
	if err != nil {
		log.Fatal("Error in portfolio File ", portfolioFile, ": ", err)
	}
	json.Unmarshal(rawJSON, &crypto.cryptoArray)

	return crypto

}

func GetRate(crypto *Cryptofolio) error {

	s := make([]string, len(crypto.cryptoArray))
	for _, c := range crypto.cryptoArray {
		s = append(s, c.Name)
	}

	ticker, err := coinmarketcap.GetData(s)
	fmt.Println(ticker)
	if err != nil {
		fmt.Println("error ticker")
		return err
	}
	for _, c := range crypto.cryptoArray {
		c.Rate = ticker.Coins[c.Name].PriceEUR
	}

	return nil
}
