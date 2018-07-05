package main

import (
	"cryptohedge"
	//"github.com/cryptohazard/hackico"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// GOOS=windows GOARCH=amd64 go build -o cryptohedge.exe main.go
type Data struct {
	Value      float64
	Index      float64
	Percentage float64
}

func main() {
	http.HandleFunc("/", helloWorld)
	fs := http.FileServer(http.Dir("../images/"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	process(w)
}

func process(w http.ResponseWriter) float64 {
	fmt.Println("\n ***Reading the portfolio!***\n")
	funds := cryptohedge.ParseJSON("portfolio.json")
	funds.Print()
	err := cryptohedge.GetRate(funds)
	if err != nil {
		fmt.Fprintf(w, "Something went wrong getting the cryptocurrencies price\n")
		fmt.Fprint(w, err)
		return 0.0
	}
	value := funds.Value()

	tmpl := template.Must(template.ParseFiles("../index.html"))
	data := Data{
		Value: roundup(value),
	}

	tmpl.Execute(w, data)

	return value

}

func roundup(f float64) float64 {
	return float64(int(f*100)) / 100
}
