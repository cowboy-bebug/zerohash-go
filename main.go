package main

import "github.com/cowboy-bebug/zerohash-go/helper"

var exit = make(chan bool)

func main() {
	for _, pair := range helper.Pairs {
		c := make(chan float64)
		p := &helper.Price{}
		go helper.Produce(c, pair)
		go helper.Consume(c, pair, p)
	}
	<-exit
}
