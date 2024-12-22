package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá Mundo!", canal)

	fmt.Println("Após início da execução da função escrever")

	/*
		for{
			mensagem, aberto := <- canal

			if !aberto {
				break
			}

			fmt.Println(mensagem)
		}
	*/

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa!")
}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)

	for i := 0; i < 5; i++ {
		canal <- texto + strconv.Itoa(i)
		time.Sleep(time.Second)
	}

	close(canal)
}
