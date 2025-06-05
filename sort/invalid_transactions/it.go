package invalid_transactions

import (
	"sort"
	"strconv"
	"strings"
)

func invalidTransactions(transactions []string) []string {
	type tx struct {
		name   string
		occur  int
		amount int
		city   string
		index  int
	}
	txs := make(map[string][]tx)
	for i, s := range transactions {
		ts := strings.Split(s, ",")

		occur, _ := strconv.Atoi(ts[1])
		amount, _ := strconv.Atoi(ts[2])
		t := tx{
			name:   ts[0],
			occur:  occur,
			amount: amount,
			city:   ts[3],
			index:  i,
		}
		txs[t.name] = append(txs[t.name], t)
	}

	ans := make([]string, 0)
	tagged := make(map[int]bool)

	addToAns := func(i int) {
		if tagged[i] {
			return
		}

		ans = append(ans, transactions[i])
		tagged[i] = true
	}

	for _, ts := range txs {
		sort.Slice(ts, func(i, j int) bool {
			return ts[i].occur < ts[j].occur
		})

		var j int
		for i := 0; i < len(ts); i++ {
			j = i + 1
			if ts[i].amount > 1000 {
				addToAns(ts[i].index)
			}

			for j < len(ts) && ts[j].occur-ts[i].occur <= 60 {
				if ts[i].city != ts[j].city {
					addToAns(ts[i].index)
					addToAns(ts[j].index)
				}
				j++
			}
		}
	}

	return ans
}
