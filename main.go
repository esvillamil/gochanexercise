// Ejercicio de Lista de numeros por canal
package main

import (
"fmt"
"math/rand"	
)

func GetRandNumbers(numSize int, c chan int) {

	for i := 0; i < numSize; i++ {
		c <- rand.Int()
	}
	close(c)
}




func main() {

	numSize := 10
	c := make(chan int)
	go GetRandNumbers(numSize, c)
	for i := range c {
		fmt.Println(i)
	}
}
