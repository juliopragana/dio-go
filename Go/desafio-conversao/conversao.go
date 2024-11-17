// Exemplo de declaração e uso de variáveis
package main

import "fmt"

//declarando uma constante
const ebulicaoK float64 = 373.15

func main() {
	var tempK float64 = ebulicaoK      //Variavel da temperaura da ebulição em graus Kelvin
	var tempC float64 = tempK - 273.15 //Variavel da temperaura da ebulição em graus Celsius

	// fmt.Println("A temperatua de eboluição da água em ºF =", tempK)
	// fmt.Println("A temperatua de eboluição da água em ºC =", tempC)
	fmt.Printf("A temperatua de eboluição da água em Kelvin é %gºK e em graus Celsiuns é de %gCº ", tempK, tempC)
}
