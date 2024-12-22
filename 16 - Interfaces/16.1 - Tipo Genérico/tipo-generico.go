package main

import "fmt"

func generica(g interface{}) {
	fmt.Println(g)
}

func main() {
	generica("String")
	generica(1234)
	generica(15.8)
	generica(true)

	fmt.Println("Oi", 5, true, 'c', float64(12.8))

	mapa_generico := map[interface{}]interface{}{
		1:               "String",
		"nome":          1234,
		true:            false,
		float32(478.63): `c`,
	}

	fmt.Println(mapa_generico)

}
