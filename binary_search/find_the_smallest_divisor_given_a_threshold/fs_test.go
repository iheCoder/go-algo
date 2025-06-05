package find_the_smallest_divisor_given_a_threshold

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestGenerateRandomArray(t *testing.T) {
	low, high := 1, 1000000
	x, y := 1, 50000
	arr := generateRandomArray(low, high, x, y)
	arrStrs := make([]string, 0, len(arr))
	for _, num := range arr {
		arrStrs = append(arrStrs, strconv.Itoa(num))
	}
	s := strings.Join(arrStrs, ",")
	fmt.Println(s)
}

func generateRandomArray(low, high, x, y int) []int {
	if low > high {
		fmt.Println("Invalid range: low must be less than or equal to high.")
		return nil
	}
	if x > y {
		fmt.Println("Invalid size range: x must be less than or equal to y.")
		return nil
	}

	rand.Seed(time.Now().UnixNano()) // 设置随机数种子

	// 随机生成数组大小
	arraySize := rand.Intn(y-x+1) + x

	// 初始化数组并填充随机数
	array := make([]int, arraySize)
	for i := 0; i < arraySize; i++ {
		array[i] = rand.Intn(high-low+1) + low
	}

	return array
}
