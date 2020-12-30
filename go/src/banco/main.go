package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta VerificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type VerificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(200)

	PagarBoleto(&contaExemplo, 60)
	fmt.Println(contaExemplo.ObterSaldo())

	contaExemplo2 := contas.ContaPoupanca{}
	contaExemplo2.Depositar(400)
	PagarBoleto(&contaExemplo2, 120)

	fmt.Println(contaExemplo2.ObterSaldo())

}
