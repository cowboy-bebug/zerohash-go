package helper

import (
	"log"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

const (
	qty       int           = 200
	wsAddr    string        = "wss://ws-feed.exchange.coinbase.com"
	wsTimeout time.Duration = 5
)

type (
	MsgSub struct {
		Type       string   `json:"type"`
		ProductIds []string `json:"product_ids"`
		Channels   []string `json:"channels"`
	}
	MsgMatch struct {
		Type      string `json:"type"`
		ProductId string `json:"product_id"`
		Price     string `json:"price"`
		Message   string `json:"message"`
	}
	Price struct {
		Prices []float64
		Sum    float64
	}
)

var Pairs = []string{
	"BTC-USD",
	"ETH-USD",
	"ETH-BTC",
}

func subscribe() (*websocket.Conn, error) {
	conn, _, err := (&websocket.Dialer{
		HandshakeTimeout: wsTimeout * time.Second,
	}).Dial(wsAddr, nil)
	err = conn.WriteJSON(MsgSub{
		Type:       "subscribe",
		ProductIds: Pairs,
		Channels:   []string{"matches"},
	})
	return conn, err
}

func Produce(c chan<- float64, pair string) {
	conn, err := subscribe()
	if err != nil {
		panic(err)
	}

	for {
		msg := &MsgMatch{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("something went wrong: %s\n", err)
			// TODO: push to DLQ for further processing / analysis
		}
		if msg.Type == "error" {
			log.Printf("something went wrong: %s\n", msg.Message)
			// TODO: push to DLQ for further processing / analysis
		}

		if msg.ProductId == pair {
			price, err := strconv.ParseFloat(msg.Price, 64)
			if err != nil {
				log.Printf("something went wrong: %s\n", err)
				// TODO: push to DLQ for further processing / analysis
			}
			c <- price
		}
	}
}

func computeVwap(price float64, p *Price) float64 {
	if len(p.Prices) >= qty {
		first, rest := p.Prices[0], p.Prices[1:]
		p.Sum -= first
		p.Prices = rest
	}
	p.Sum += price
	p.Prices = append(p.Prices, price)
	return p.Sum / float64(len(p.Prices))
}

func Consume(c <-chan float64, pair string, p *Price) {
	for {
		price := <-c
		vwap := computeVwap(price, p)
		log.Printf("[%s]: vwap: %f\n", pair, vwap)
	}
}
