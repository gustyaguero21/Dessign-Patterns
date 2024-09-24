package main

import "fmt"

// PaymentStrategy es una interfaz que define el método de pago que deben implementar todas las estrategias de pago.
type PaymentStrategy interface {
	Payment(amount float64) string // Este método es implementado por todas las estrategias de pago y recibe un monto a pagar.
}

// CreditCard representa una estructura para el método de pago con tarjeta de crédito.
type CreditCard struct {
	cardNumber string // Número de la tarjeta de crédito.
}

// Payment realiza un pago con tarjeta de crédito y retorna un mensaje confirmando el pago.
func (c *CreditCard) Payment(amount float64) string {
	return fmt.Sprintf("Ha pagado %f con la tarjeta de número %s", amount, c.cardNumber)
}

// PayPal representa una estructura para el método de pago con PayPal.
type PayPal struct {
	email string // Correo electrónico asociado a la cuenta de PayPal.
}

// Payment realiza un pago con PayPal y retorna un mensaje confirmando el pago.
func (pp *PayPal) Payment(amount float64) string {
	return fmt.Sprintf("Ha pagado %f con la cuenta %s", amount, pp.email)
}

// Crypto representa una estructura para el método de pago con criptomonedas.
type Crypto struct {
	walletAddress string // Dirección de la billetera de criptomonedas.
}

// Payment realiza un pago con criptomonedas y retorna un mensaje confirmando el pago.
func (c *Crypto) Payment(amount float64) string {
	return fmt.Sprintf("Ha pagado %f con la wallet registrada en %s", amount, c.walletAddress)
}

// PaymentProcessor es una estructura que gestiona el proceso de pago y permite seleccionar una estrategia de pago.
type PaymentProcessor struct {
	strategy PaymentStrategy // Estrategia de pago seleccionada.
}

// SetPaymentStrategy permite establecer el método de pago que se usará en el procesador de pagos.
func (p *PaymentProcessor) SetPaymentStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

// Pay realiza el pago utilizando la estrategia seleccionada y muestra un mensaje con los detalles del pago.
func (p *PaymentProcessor) Pay(amount float64) {
	fmt.Println(p.strategy.Payment(amount)) // Llama al método Payment de la estrategia seleccionada.
}

// Función principal
// func main() {
// 	newPayment := PayPal{
// 		email: "gustyaguero21@gmail.com", // Crea una nueva instancia de PayPal con un email asociado.
// 	}

// 	strategy := PaymentProcessor{} // Crea una nueva instancia del procesador de pagos.

// 	strategy.SetPaymentStrategy(&newPayment) // Establece la estrategia de pago a PayPal.

// 	strategy.Pay(10000) // Realiza un pago de 10,000 utilizando la estrategia seleccionada.
// }
