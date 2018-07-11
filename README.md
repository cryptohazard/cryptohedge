# Cryptohedge = cryptocurrency + porfolio + hedge fund
This project is about a cryptocurrency hedge fund. It lets you visualize and analyzes the current state of your portfolio. It is not a wallet so does not deal with private or public keys. No funds are at stake if the server you host it on is hacked. This is also useful if you do not want to trust a third party with your data.
 The goal is to build various tools to:
* visualize the health of the cryptofolio
* manage the cryptofolio
* analyze the, past and present, portfolios: performance, diversification, correlation to various data...
* have access controls for users(https, login)

## Installation and compilation
You need to have a working Go(lang) environment in version >=1.9 and clone this repository in your ```src``` directory.

Now you can go in the ```exec``` folder and build the executable( I added the ```-o``` option to change the name):
```
$ cd exec/
$ go build -o cryptohedge main.go
$ sudo ./cryptohedge
```
The port is ```8080```.

Note: You can build for  windows using:
```
GOOS=windows GOARCH=amd64 go build -o cryptohedge.exe main.go
```

## configuration
For now, the **cryptohedge** data are submitted through JSON file.
The file ```portfolio.json```  is  for the cryptocurrencies names and their amount. The name should be the one used by [coinmarketcap](http://coinmarketcap.com/), the price provider we use. For instance Bitcoin is ```bitcoin``` while Bitcoin Cash is ```bitcoin-cash```.

The portfolio file has the following format:
```
[
  {
    "name" : "ethereum",
	   "amount": 1.1
   },
   {
     "name" : "bitcoin",
 	   "amount": 2.2
    },
    ...
]
```

The file ```shares.json``` is for the users names and the number of shares in the fund they have. The format is the following:
```
[
  {
    "name" : "aba",
	   "shares": 18
   },
   {
     "name" : "aora",
 	   "shares": 26
    },
    ...
]
```
Example files, with obviously fake data, are provided.

## Features
### "/"
This shows the value(in €) of the cryptofolio. It also shows the current state of the market, courtesy of [coin360](https://coin360.io/).

### "/cryptofolio"
This shows the content of the portfolio. The coin name and amount are retrieved from the ```portfolio.json```. The value is computed from the price available through [coinmarketcap API](https://coinmarketcap.com/api/). We also compute the percentage of each coin in the portfolio. This is interesting to see how dependent you are of a particular currency. Next versions will allow you to set your *dream portfolio* so that you can compare it to the current one.

### "/cryptohedge"
Here it is mainly about users. How many shares do they *HODL*? What the value in €? It also shows the *Index*, the price of share in the hedge fun. This page should be locked to *administrators* in the future and users will only see their own info.

## Frontend development
The Frontend part is clearly not *my forte*. As such, I am just doing quick-and-dirty pages that shows what can be done with value retrieved from a Golang Backend. I invite you to check [Golang html templates package](https://golang.org/pkg/html/template/) to understand the relation between the frontend and the backend but they can be developped separately.
I plan to make a task request on Utopian.io for the Frontend and a logo.

## Tentative Roadmap
* add fiat support so that you can still track your portfolio when you have to convert back and forth from fiat(*tether* is not fiat)
* compare current portfolio with ideal/past/different portfolio
* add  database and/or CSV support to do more complex analysis. For instance it would be interesting to see if the weekly performances are correlated, if portfolio is *hedged enough*, if other statistical indicators would be interesting...
* more hedge fund features: entry amount, overall performance of the hedge fund, historical variation, users personalized history
* TLS support
* blog = Steem articles rendering
