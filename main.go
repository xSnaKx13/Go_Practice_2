package main

import (
	"errors"
	"fmt"
)

func main(){
	var transactions[] float64 
	for{
		fmt.Print("Введите транзакцию: ")
		transactionN, err := newTransaction()
		if err!=nil{
			fmt.Println("Ошибка: ", err)
			continue
		}
		if transactionN == 0{
			break
		}
		transactions = append(transactions, transactionN)
	}
	fmt.Println(transactions)
}

func newTransaction() (float64, error){
	var transaction float64 
	var err error
	_, err = fmt.Scan(&transaction)
	if err!=nil{
		return 0, errors.New("Не является транзакцией")
	}
	return transaction, nil
}
