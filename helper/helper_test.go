package helper

import (
	"math/rand"
	"testing"
)

func createPrice(n int) ([]float64, float64, Price) {
	var prices []float64
	var sum float64

	for i := 0; i <= n; i++ {
		price := rand.Float64()
		prices = append(prices, price)
		sum += price
	}

	p := Price{
		Prices: prices,
		Sum:    sum,
	}

	return prices, sum, p
}

func TestComputeVwap(t *testing.T) {
	_, sum, p := createPrice(10)
	price := rand.Float64()
	vwap := computeVwap(price, &p)
	expected := (sum + price) / float64(len(p.Prices))

	if vwap != expected {
		t.Errorf("[ComputeVwap] got: %f, want: %f.", vwap, expected)
	}
}

func TestComputeVwapForFullSlice(t *testing.T) {
	prices, sum, p := createPrice(qty)
	price := rand.Float64()
	vwap := computeVwap(price, &p)
	expected := (sum - prices[0] + price) / float64(len(p.Prices))

	if vwap != expected {
		t.Errorf("[ComputeVwap] got: %f, want: %f.", vwap, expected)
	}
}
