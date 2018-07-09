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
	http.HandleFunc("/cryptofolio", cryptofolio)
	fs := http.FileServer(http.Dir("../images/"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))
	fs = http.FileServer(http.Dir("../css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	fs = http.FileServer(http.Dir("../js/"))
	http.Handle("/js/", http.StripPrefix("/js/", fs))
	fs = http.FileServer(http.Dir("../vendor/"))
	http.Handle("/vendor/", http.StripPrefix("/vendor/", fs))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	funds, err := process(w)
	if err != nil {
		return
	}

	value := funds.Value()

	tmpl := template.Must(template.ParseFiles("../index.html"))
	data := Data{
		Value: roundup(value),
	}
	tmpl.Execute(w, data)
}

func cryptofolio(w http.ResponseWriter, r *http.Request) {
	funds, err := process(w)
	if err != nil {
		return
	}

	tmpl := template.Must(template.ParseFiles("../cryptofolio.html"))

	tmpl.Execute(w, funds)
}
func process(w http.ResponseWriter) (*cryptohedge.Cryptofolio, error) {
	fmt.Println("\n ***Reading the portfolio!***\n")
	funds := cryptohedge.ParseJSON("portfolio.json")
	funds.Print()
	err := cryptohedge.GetRate(funds)
	if err != nil {
		message := "Something went wrong getting the cryptocurrencies price\n"
		pageNotFound(w, message)
		fmt.Println(err)
		return new(cryptohedge.Cryptofolio), err
	}
	return funds, err
}

func roundup(f float64) float64 {
	return float64(int(f*100)) / 100
}

func pageNotFound(w http.ResponseWriter, message string) {
	tmpl := template.Must(template.ParseFiles("../404.html"))
	data := struct{ Message string }{Message: message}
	tmpl.Execute(w, data)
}
