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

type Cryptohedge struct {
	Index      float64
	Total      float64
	ShareArray []*Share
}

type Share struct {
	Name   string  `json:"name"`
	Shares float64 `json:"shares"`
	Value  float64
}

func (c *Coin) computeValue() (value float64) {
	value = c.Amount * c.Rate
	c.Value = value
	return
}

func (c *Coin) computePercentage(value float64) (p float64) {
	p = percentage(c.Value, value)
	c.Percentage = p
	return
}

func (crypto *Cryptofolio) Value() (value float64) {
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

func (hedge *Cryptohedge) ComputeValues(value float64) {
	for _, s := range hedge.ShareArray {
		hedge.Total += s.Shares
	}
	hedge.Index = value / hedge.Total
	for _, s := range hedge.ShareArray {
		s.Value = hedge.Index * s.Shares
	}
}

func (hedge *Cryptohedge) Print() {
	for _, s := range hedge.ShareArray {
		fmt.Println(s.Name, " ", s.Shares, " ", s.Value)
	}
}
