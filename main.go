package main

import (
	"basic-go-package/calculator"
	"basic-go-package/persegi"
	"fmt"
)

func main() {
	fmt.Println("Test")

	var c calculator.HitungCalculator
	c = calculator.Nilai{4, 7, 10}
	fmt.Println(c.Perkalian())

	var p persegi.HitungPersegi
	p = persegi.Persegi{90}
	fmt.Println(p.Luas())
}
