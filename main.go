package main

import (
	"fmt"
)

func main()  {
	const eurToUsd = 1.2
	const usdToRub = 90.0
	const eurToRub = eurToUsd * usdToRub

}
func inputUserValue(usingValute, convertValue float64){
	fmt.Println("Введите валюту которая у вас есть")
	fmt.Scan(&usingValute)
	fmt.Println("Введите валюту на которую необходимо обменять")
	fmt.Scan(&convertValue)
}
func convert (usingValute, convertValue, howMatch float64) float64{
	
}
