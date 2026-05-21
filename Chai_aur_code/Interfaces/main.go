package main

import "fmt"

type payment struct{}

func (p payment) makePayment(amount float32){

}

type rezorpay struct{

}

func (r rezorpay) pay(amount float32){
	//logic to make payment 
	fmt.Println("making payment using rezorpay",amount)
}

func main(){
 

}