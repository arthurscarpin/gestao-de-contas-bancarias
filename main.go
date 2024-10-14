package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}
	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "O depósito foi realizado com sucesso", c.saldo
	}
	return "Valor do depósito menor que zero", c.saldo
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)
	// fmt.Println(contaDaSilvia.Sacar(400))
	status, valor := contaDaSilvia.Depositar(400)
	fmt.Println(status, valor)
}
