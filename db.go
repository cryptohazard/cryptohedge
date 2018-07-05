package cryptohedge

// before we work with a db, we are going to store everything in a struct
import (
	"fmt"
)

type Cryptofolio struct {
	fiatArray   []*Coin
	cryptoArray []*Coin
}

type Coin struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Rate   float64
}

func (c *Coin) Value() (value float64) {
	value = c.Amount * c.Rate
	fmt.Println(c.Name, " ", value)
	return
}

func (crypto *Cryptofolio) Value() (value float64) {
	fmt.Println("\n***Coins value***\n")
	for _, c := range crypto.cryptoArray {
		value += c.Value()
	}

	return
}

func (crypto *Cryptofolio) Print() {
	for _, c := range crypto.cryptoArray {
		fmt.Println(c.Name, " ", c.Amount)
	}
}
