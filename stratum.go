
package main

import (
	"github.com/zond/gomarket"
)

func main() {
	m := gomarket.NewMarket()
	martin := gomarket.NewSimpleActor("martin")
	rice := "rice"
	martin.Bid(10.0, rice, 2.0)
	emelie := gomarket.NewSimpleActor("emelie")
	emelie.Ask(20.0, rice, 1.5)
	m.Actors[martin] = true
	m.Actors[emelie] = true
	m.Trade()
}