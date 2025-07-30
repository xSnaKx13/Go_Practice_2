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
			fmt.Println(err)
		}
		if transactionN == 0{
			break
		}
		transactions = append(transactions, transactionN)
	}
	fmt.Print(transactions)
}

func newTransaction()(float64, error){
	var transaction float64 
	var err error
	_, err = fmt.Scan(&transaction)
	if err!=nil{
		errors.New("Не является транзакцией")
		return 0, err
	}
	return transaction, nil
}
