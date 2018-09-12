//Downtiser
package main

import (
	"fmt"
	"math"
	"strconv"
)

func prime(n int) bool {
	for i:=2;float64(i)<=math.Sqrt(float64(n));i++  {
		if (n%i == 0) {
			return false
		}
	}
	return true
}

func mypow(m int, n int) int {
	pow := 1
	for i:=1;i<=n;i++  {
		pow *= m
	}
	return pow
}

func narcissus(n int)  {
	one := n % 10
	ten := n % 100 / 10
	hundred := n/100
    sum := mypow(one, 3) + mypow(ten, 3) + mypow(hundred, 3)
	if sum == n {
		fmt.Println(n)
	}
}

func newNarcissus(str string)  {
	var sum int
	length := len(str)
	for i:=0;i<length;i++  {
		pow := int(str[i] - '0')
		sum += mypow(pow, 3)
	}
	n, err := strconv.Atoi(str)
	if err != nil {
		println(err)
		return
	}
	if sum == n {
		fmt.Printf("narcissus:[%s] \n", str)
	}
}

func factorial(n int) int {
	res := 1
	sum := 1
	for i:=2;i<=n;i++  {
		res *= i
		sum += res
	}
	return sum
}

func main()  {
	count := 0
	for i:=101;i<200;i++  {
		if prime(i) {
			count++
			fmt.Println(i)
		}
	}
	fmt.Println("count:", count)
	for i:=100;i<1000;i++ {
		newNarcissus(fmt.Sprintf("%d", i))
	}
	var n int
	fmt.Printf("please input>>>")
	fmt.Scanf("%d", &n)
	fmt.Println("The factorial:", factorial(n))
}