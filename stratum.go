
package main

import (
	"../gomarket"
	"fmt"
)

func main() {
	m := gomarket.NewMarket()
	martin := gomarket.NewSimpleActor("martin")
	elise := gomarket.NewSimpleActor("elise")
	emelie := gomarket.NewSimpleActor("emelie")
	mattias := gomarket.NewSimpleActor("mattias")
	rice := "rice"
	shoes := "shoes"
	pizza := "pizza"
	hats := "hats"
	martin.Ask(10.0, hats, 1.0)
	emelie.Ask(20.0, hats, 1.2)
	elise.Ask(15.0, hats, 0.8)
	mattias.Bid(17, hats, 10.0)
	martin.Bid(10.0, rice, 2.0)
	elise.Bid(12.0, rice, 2.1)
	mattias.Bid(8.0, rice, 0.5)
	emelie.Ask(20.0, rice, 1.5)
	emelie.Ask(2.0, shoes, 20.0)
	martin.Bid(2.0, pizza, 50.0)
	mattias.Ask(1.0, shoes, 15.0)
	m.Actors[martin] = true
	m.Actors[emelie] = true
	m.Actors[mattias] = true
	m.Actors[elise] = true
	m.Trade()
	fmt.Println("prices: ", m.Prices)
}