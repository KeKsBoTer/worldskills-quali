package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
)

func forEachLine(filepath string, fn func(int)) {
	b, _ := ioutil.ReadFile(filepath)

	for _, line := range strings.Split(string(b), "\r\n") {
		num, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		fn(num)
	}
}

func main() {
	count := 0
	min, max := 0, 0
	sum := 0
	product := big.NewInt(1)
	frequency := map[int]int{}
	forEachLine("Task_01_Data.txt", func(num int) {
		if num >= 2023 {
			count++
		}
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
		sum += num
		product.Mul(product, big.NewInt(int64(num)))
		frequency[num]++
	})
	fmt.Println("Creater than 2023:", count)
	fmt.Println("Diff. between min and max: ", max-min)
	fmt.Println("Sum:", sum)
	prod := product.String()
	fmt.Println("Product (last 4 digits):", prod[len(prod)-4:])

	fmt.Println("Frequency:")
	for _, n := range []int{1115, 7139, 4347, 7847, 9159, 3921, 1997, 3755, 8415} {
		fmt.Printf("	%d: %d\n", n, frequency[n])
	}
}
