package cryptohedge

// before we work with a db, we are going to store everything in a struct
import (
	"fmt"
)

type Cryptofolio struct {
	FiatArray   []*Coin
	CryptoArray []*Coin
}

type Coin struct {
	Name       string  `json:"name"`
	Amount     float64 `json:"amount"`
	Rate       float64
	Percentage float64
	Value      float64
}

func (c *Coin) computeValue() (value float64) {
	value = c.Amount * c.Rate
	c.Value = value
	//fmt.Println(c.Name, " ", value)
	return
}

func (c *Coin) computePercentage(value float64) (p float64) {
	p = percentage(c.Value, value)
	c.Percentage = p
	return
}

func (crypto *Cryptofolio) Value() (value float64) {
	//fmt.Println("\n***Coins value***\n")
	for _, c := range crypto.CryptoArray {
		value += c.computeValue()
	}
	return
}

func (crypto *Cryptofolio) Percentage() {
	value := crypto.Value()
	for _, c := range crypto.CryptoArray {
		c.computePercentage(value)
	}

}

func (crypto *Cryptofolio) Print() {
	for _, c := range crypto.CryptoArray {
		fmt.Println(c.Name, " ", c.Amount, " ", c.Value, " ", c.Percentage, "%")
	}
}

// compute percentage
func percentage(part float64, total float64) (p float64) {
	p = 100 * part / total
	return
}
