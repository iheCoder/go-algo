package k_closest_points_to_origin

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestKC(t *testing.T) {

}

func TestGenTestCases(t *testing.T) {
	n := 10000
	printCoordinates(generateCoordinates(n))
}

func generateCoordinates(n int) [][]int {
	coordinates := make([][]int, n)
	for i := 0; i < n; i++ {
		// 示例：生成随机坐标，这里用i和i*2代替
		x := rand.Intn(10000)
		y := rand.Intn(10000)
		isNeg := rand.Int()%2 == 0
		if isNeg {
			x = -x
		}
		isNeg = rand.Int()%3 == 0
		if isNeg {
			y = -y
		}
		coordinates[i] = []int{x, y}
	}
	return coordinates
}

func printCoordinates(coordinates [][]int) {
	fmt.Print("[")
	for i, coord := range coordinates {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("[%d, %d]", coord[0], coord[1])
	}
	fmt.Println("]")
}
