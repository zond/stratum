
package main

import (
	"../gomarket"
	"fmt"
)

type SimpleActor struct {
	name string
	asks map[*gomarket.Order]bool
	bids map[*gomarket.Order]bool
}
func NewSimpleActor(name string) *SimpleActor {
	return &SimpleActor{name, make(map[*gomarket.Order]bool), make(map[*gomarket.Order]bool)}
}
func (a *SimpleActor) String() string {
	return a.name
}
func (a *SimpleActor) Ask(units float64, resource interface{}, price float64) {
	a.asks[&gomarket.Order{units, resource, price, a}] = true
}
func (a *SimpleActor) Bid(units float64, resource interface{}, price float64) {
	a.bids[&gomarket.Order{units, resource, price, a}] = true
}
func (a *SimpleActor) Asks() map[*gomarket.Order]bool {
	return a.asks
}
func (a *SimpleActor) Bids() map[*gomarket.Order]bool {
	return a.bids
}
func (a *SimpleActor) Buy(bid, ask *gomarket.Order, price float64) {
	fmt.Println(a, "buys", ask.Units, ask.Resource, "á", price, "from", ask.Actor)
	ask.Actor.Deliver(bid, ask, price)
}
func (a *SimpleActor) Deliver(bid, ask *gomarket.Order, price float64) {
	fmt.Println(a, "delivers", bid.Units, bid.Resource, "á", price, "to", bid.Actor)
}

func main() {
	m := gomarket.NewMarket()
	martin := NewSimpleActor("martin")
	elise := NewSimpleActor("elise")
	emelie := NewSimpleActor("emelie")
	mattias := NewSimpleActor("mattias")
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