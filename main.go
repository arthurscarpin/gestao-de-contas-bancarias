package main

import (
	"fmt"
	"gestao-de-contas-bancarias/contas"
)

func main() {
	contaDoDenis := contas.ContaPoupanca {}
	contaDoDenis.Depositar(100)
	contaDoDenis.Sacar(555)
	fmt.Println(contaDoDenis.ObterSaldo())
}
