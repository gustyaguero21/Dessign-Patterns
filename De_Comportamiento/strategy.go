package main

import "fmt"

type PaymentStrategy interface { //primero definimos la interfaz de pagos. Ella se encargara de tomar el pago y
	Payment(amount float64) string //este metodo lo implementan todos los metodos de pago ya que contiene su firma.
}

type CreditCard struct { //estructura de la tarjeta de credito
	cardNumber string
}

func (c *CreditCard) Payment(amount float64) string { //realiza el pago con tarjeta de credito.
	return fmt.Sprintf("Ha pagado %f con la tarjeta de numero %s", amount, c.cardNumber)
}

type PayPal struct { //estructura de paypal
	email string
}

func (pp *PayPal) Payment(amount float64) string { //realiza el pago con paypal.
	return fmt.Sprintf("Ha pagado %f con la cuenta %s ", amount, pp.email)
}

type Crypto struct { //estructura de criptomonedas.
	walletAddress string
}

func (c *Crypto) Payment(amount float64) string { //realiza el pago con criptomonedas.
	return fmt.Sprintf("Ha pagado %f con la wallet registrada en %s", amount, c.walletAddress)
}

type PaymentProcessor struct { //estructura que aloja todos los medios de pago ya que tiene la firma de paymentstrategy.
	strategy PaymentStrategy
}

func (p *PaymentProcessor) SetPaymentStrategy(strategy PaymentStrategy) { //aqui se elige el metodo de pago y el metodo payment correspondiente al metodo de pago seleccionado.
	p.strategy = strategy
}

func (p *PaymentProcessor) Pay(amount float64) { //una vez seleccionado el metodo de pago, solo recibe un monto e implementa el pago correspondiente.
	fmt.Println(p.strategy.Payment(amount))
}

func main() {
	newPayment := PayPal{
		email: "gustyaguero21@gmail.com",
	}

	strategy := PaymentProcessor{}

	strategy.SetPaymentStrategy(&newPayment)

	strategy.Pay(10000)

}
