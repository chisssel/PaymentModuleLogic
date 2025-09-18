package main

import (
	"PaymentModuleLogic/payments"
	"PaymentModuleLogic/payments/methods"
	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewCrypto()

	paymentModule := payments.NewPaymentModule(method)

	paymentModule.Pay("Бургер", 5)
	idPhone := paymentModule.Pay("Телефон", 500)
	idGame := paymentModule.Pay("Игра", 20)

	paymentModule.Cancel(idPhone)

	allInfo := paymentModule.AllInfo()
	pp.Println(allInfo)

	gameInfo := paymentModule.Info(idGame)
	pp.Println("Game info:", gameInfo)
}
