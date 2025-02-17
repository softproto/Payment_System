package main

import (
	"fmt"
	"payment_system/pkg/PaymentSystem"
)

// демонстрация работы пакета PaymentSystem
func main() {

	//Создаем плfтежную систему BelPay
	BelPay := PaymentSystem.New()

	//Эмиссионный счет в BYN
	fmt.Println("Эмиссионный счет в BYN: ", PaymentSystem.EmitionAccount("BYN"))

	// Счет для вывод денежных средств BYN из оборота
	fmt.Println("Счет для вывод денежных средств BYN из оборота", PaymentSystem.WithdrowalAccount("BYN"))

	//Эмиссия денежных средств - 1000 BYN
	err := BelPay.Emition("BYN", 1000)
	if err != nil {
		fmt.Println("Ошибка эмиссии: ", err)
	}

	//Эмиссия денежных средств - 1000 USD
	err = BelPay.Emition("USD", 1000)
	if err != nil {
		fmt.Println("Ошибка эмиссии: ", err)
	}

	//Создаем аккаунт на имя "MAZ" в валюте "BYN"
	err = BelPay.CreateAccount("BY13NBRB3600900000002Z00AB22", "MAZ", "BYN", "active")
	if err != nil {
		fmt.Println("Ошибка создания аккаунта: ", err)
	}

	//Кредитуем "MAZ" на 100 BYN (переводим с эмиссионного счета 100 BYN)
	err = BelPay.Transfer(PaymentSystem.EmissionAccounts["BYN"], "BY13NBRB3600900000002Z00AB22", 100)
	if err != nil {
		fmt.Println("Ошибка оплаты:", err)
	}

	//Создаем аккаунт на имя "EPAM" в валюте "USD"
	err = BelPay.CreateAccount("BY13NBRB3600900000002Z00AB33", "EPAM", "USD", "active")
	if err != nil {
		fmt.Println("Ошибка создания аккаунта: ", err)
	}

	//Кредитуем "EPAM" на 100 USD (переводим с эмиссионного счета 100 USD)
	BelPay.Transfer(PaymentSystem.EmissionAccounts["BYN"], "BY13NBRB3600900000002Z00AB33", 100)
	if err != nil {
		fmt.Println("Ошибка оплаты:", err)
	}

	//Создаем аккаунт на имя "EPAM" в валюте "BYN"
	err = BelPay.CreateAccount("BY13NBRB3600900000002Z00AB44", "EPAM", "BYN", "active")
	if err != nil {
		fmt.Println("Ошибка создания аккаунта: ", err)
	}

	//"MAZ" совершает платеж в пользу "EPAM" 10 BYN
	err = BelPay.Transfer("BY13NBRB3600900000002Z00AB22", "BY13NBRB3600900000002Z00AB44", 10)
	if err != nil {
		fmt.Println("Ошибка оплаты:", err)
	}

	//"MAZ" совершает платеж в пользу "EPAM" 40 BYN (данные передаются в json формате)
	j := []byte(`{"sender":"BY13NBRB3600900000002Z00AB22", "recipient":"BY13NBRB3600900000002Z00AB44", "amount":20}`)
	err = BelPay.Payment(j)
	if err != nil {
		fmt.Println("Ошибка оплаты:", err)
	}

	err = BelPay.Withdrowal("BY13NBRB3600900000002Z00AB22", 30)
	if err != nil {
		fmt.Println("Ошибка вывода из оборота: ", err)
	}

	//Баланс счета "BY13NBRB3600900000002Z00AB22" после всех платежей
	fmt.Println("Текущий баланс BY13NBRB3600900000002Z00AB22 ", BelPay.Name("BY13NBRB3600900000002Z00AB22"), BelPay.Balance("BY13NBRB3600900000002Z00AB22"))

	//Список всех открытых счетов в системе
	a, err := BelPay.List()
	if err != nil {
		fmt.Println("Ошибка генерации списка :", err)
	} else {
		fmt.Println("Список открытых счетов: ")
		fmt.Println(string(a))
	}

}
