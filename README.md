# cryptohedge
This project is a website to track the value of your crypto-portfolio with managing a hedge fund in mind. 

## Installation and compilation
You need to have a working Go(lang) environment in version >=1.9 and clone this repository in your ```src``` directory. 

Now you can go in the ```exec``` folder and build the executable( I added the ```-o``` option to change the name):
```
$ cd exec/
$ go build -o cryptohedge main.go
$ sudo ./cryptohedge
```
The port is ```80```. You can change it if you don't want to type ```sudo``` all the time. 

Note: You can build for  windows using:
```
GOOS=windows GOARCH=amd64 go build -o cryptohedge.exe main.go
```

## Portfolio
The portfolio file has the following format:
```
quantity1 coin1
quantity2 coin2
...
```
A file example is provided. The name should be the one used by [coinmarketcap](http://coinmarketcap.com/), the price provider we use. For instance Bitcoin is ```bitcoin``` while Bitcoin Cash is ```bitcoin-cash```.

Note: There is no need to restart the program if you change the portfolio file.

# Frontend
You can put your own index.html. Just remember to write ```{{.Value}}``` where you want the value of your portfolio computed by the backend to appear. 

# Roadmap
* add fiat support so that you can still track your portfolio when you have to convert back and forth from fiat
* compare current protfolio with ideal portfoli
* add a database to do more complex analysis
* hedge fund features: funds divided in parts, accounts balance(X parts), price of a part
* TLS support
