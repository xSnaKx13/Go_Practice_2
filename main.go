package main

import (
	"fmt"
)

func main()  {
	const eurToUsd = 1.2
	const usdToRub = 90.0
	const eurToRub = eurToUsd * usdToRub

	fmt.Println(eurToRub)
}
