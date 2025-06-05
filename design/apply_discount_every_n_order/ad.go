package apply_discount_every_n_order

type Cashier struct {
	n, count int
	discount float64
	prices   map[int]float64
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	ps := make(map[int]float64)
	for i, product := range products {
		ps[product] = float64(prices[i])
	}

	return Cashier{
		n:        n,
		count:    0,
		discount: float64(discount) / 100.0,
		prices:   ps,
	}
}

func (c *Cashier) GetBill(product []int, amount []int) float64 {
	c.count++
	var price float64
	for i, p := range product {
		price += c.prices[p] * float64(amount[i])
	}

	if c.count%c.n == 0 {
		price -= price * c.discount
	}

	return price
}
