package main

import (
	"payment_system/pkg/PaymentSystem"
)

func main() {

	//Создаем плтежную систему BelPay
	BelPay := PaymentSystem.New()
	//Эмиссия денежных средств. 1000 BYN
	BelPay.Emition("BYN", 1000)
	//Эмиссия денежных средств. 1000 USD
	BelPay.Emition("USD", 1000)
	//Создаем аккаунт на имя "MAZ" в валюте "BYN"
	BelPay.CreateAccount("BY13NBRB3600900000002Z00AB22", "MAZ", "BYN", "active")
	//Кредитуем "MAZ" на 100 BYN (переводим с эмиссионного счета 100 BYN)
	BelPay.Transfer(PaymentSystem.EmissionAccounts["BYN"], "BY13NBRB3600900000002Z00AB22", 100)

}
