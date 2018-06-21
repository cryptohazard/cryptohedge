package cryptohedge

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cryptohazard/coinmarketcap"
)

func Parse(crypto *Cryptofolio, line string) {
	lineElements := strings.Split(line, " ")
	fmt.Println(lineElements[1], " \t\t\t", lineElements[0])
	var coin = new(Coin)
	coin.name = lineElements[1]
	coin.amount, _ = strconv.ParseFloat(lineElements[0], 64)
	crypto.cryptoArray = append(crypto.cryptoArray, coin)
}

func GetRate(crypto *Cryptofolio) error {
	s := make([]string, len(crypto.cryptoArray))
	for _, c := range crypto.cryptoArray {
		s = append(s, c.name)
	}

	ticker, err := coinmarketcap.GetData(s)
	if err != nil {
		fmt.Println("error ticker")
		return err
	}
	for _, c := range crypto.cryptoArray {
		c.rate = ticker.Coins[c.name].PriceEUR
	}

	return nil
}
