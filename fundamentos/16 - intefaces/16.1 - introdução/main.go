package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A Área da forma é %0.2f\n", f.area())
}

func (r retangulo) area() float64 { //mesmo nome do método dentro da interface
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

type ValorAleatorio struct {
	valor string
}

func (v ValorAleatorio) mostrarValor() string {
	return v.valor
}

func test(interf interface {
	mostrarValor() string
}) {
	fmt.Println(interf.mostrarValor())
}

func main() {
	fmt.Println("Interfaces permites flexibilidade com tipos de dados")
	fmt.Println("Interfaces só possuem assinaturas de métodos")

	r := retangulo{10, 15}
	result := r.area()
	escreverArea(r)
	c := circulo{15}.area()
	// escreverArea(c)
	fmt.Println(result)
	fmt.Println(c)

	v:= ValorAleatorio{"Hello, World!"}
	test(v)

}
