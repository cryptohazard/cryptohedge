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
	name   string
	amount float64
	rate   float64
}

func (c *Coin) Value() (value float64) {
	value = c.amount * c.rate
	fmt.Println(c.name, " ", value)
	return
}

func (crypto *Cryptofolio) Value() (value float64) {
	fmt.Println("\n***Coins value***\n")
	for _, c := range crypto.cryptoArray {
		value += c.Value()
	}

	return
}
