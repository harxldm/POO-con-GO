package main

import "fmt"

/*type Greeter interface {
	Greet()
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hola soy %s", p.Name)
}

type text string

func (t text) Greet() {
	fmt.Printf("Hola soy %s", t)
}

func Greetall( gs ...Greeter){
	for _, g := range gs {
		g.Greet()
		fmt.Printf("\t Valor : %v, Tipo %T", g, g)
	}
}*/

type metodoDePago interface {
	pagar()
}

type paypal struct{}

func (p paypal) pagar() {
	fmt.Println("pagado por paypal")
}

type Efectivo struct{}

func (e Efectivo) pagar() {
	fmt.Println("pagado por paypal")
}

type Tarjeta struct{}

func (t Tarjeta) pagar() {
	fmt.Println("pagado por paypal")
}

func factory(metodo uint) metodoDePago {
	switch metodo {
	case 1:
		return paypal{}
	case 2:
		return Efectivo{}
	case 3:
		return Tarjeta{}
	default:
		return nil

	}
}

func main() {
	/*p := Person{Name: "Harold"}
	var t  text = "Harold"
	Greetall(p, t)*/
	var metodo uint
	fmt.Println("Dijite un numero de metodo de pago: ")
	fmt.Println("1: Paypal")
	fmt.Println("2: Efectivo")
	fmt.Println("3: Tarjeta")
	_, err := fmt.Scanln(&metodo)
	if err != nil {
		panic("Debe digitar un metodo valido")
	}

	if metodo > 3 {
		panic("Su respuesta fue mayor a 3")
	}
	f := factory(metodo)
	f.pagar()
}
