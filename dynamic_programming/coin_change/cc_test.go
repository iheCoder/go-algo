package coin_change

import "testing"

func TestCC(t *testing.T) {
	coins := []int{1, 2, 5}
	//t.Log(coinChange(coins, 11))
	//coins = []int{2}
	//t.Log(coinChange(coins, 3))
	//coins = []int{2}
	//t.Log(coinChange(coins, 0))
	//coins = []int{2}
	//t.Log(coinChange(coins, 2))
	coins = []int{186, 419, 83, 408}
	t.Log(coinChange(coins, 6249))
}
