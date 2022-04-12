//Fiquei responsável de desenvolver uma aplicação para descobrir qual seria o novo salário de um funcionário,
//baseado no aumento que ele recebeu. Esses valores são passados no formato de string, já que chegam de um formulário HTML.
//Para saber o aumento precisamos multiplicar o salário pelo aumento(%) e fazer a divisão por 100,
//depois somamos esse valor obtido ao salário atual e temos o valor do novo salário.

package main

import (
	"fmt"
	"strconv" //Conversões são feitas pelo pacote strconv.
)

func main() {
	salario := "1000"
	aumento := "20"

	//converter um tipo string para um tipo int, usamos o método Atoi do pacote strconv
	salarioConvertido, err := strconv.Atoi(salario)
	if err != nil {
		fmt.Println(err)
	}
	aumentoConvertido, err := strconv.Atoi(aumento)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("O salário é", salario, "e o aumento de", aumento+"%")

	novoSalario := ((salarioConvertido * aumentoConvertido) / 100) + salarioConvertido
	fmt.Println("O novo salário é:", novoSalario)
}

//Para saber mais

//Caso o valor do aumento ou do salário seja um valor decimal, por exemplo 22.2 ou 2000.50,
//podemos utilizar a função strconv.Itoa do mesmo pacote, para converter uma string para número decimal,
//conforme o exemplo abaixo:

// salarioConvertido, err := strconv.ParseInt(salario, 64)
// if err != nil {
// 	fmt.Println(err)
// }
// aumentoConvertido, err := strconv.ParseInt(aumento, 64)
// if err != nil {
// 	fmt.Println(err)
// }

//Mas, e o contrário?

//Agora que sabemos o valor do novo salário, precisamos retornar esse valor para o HTML no formato string.
//Nesse caso, utilizamos outra função do mesmo pacote chamada strconv.Itoa, conforme o exemplo abaixo:

// package main

// import(
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	novoSalarioInt := 2019
// 	salarioConvertido := strconv.Itoa(novoSalarioInt)
// 	fmt.Println(salarioConvertido)
// }
