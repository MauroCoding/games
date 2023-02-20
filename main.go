package main

import (
	"fmt"
	"games"
)

func main() {

	GetClases := func(c Course) string {
		txt := "Las clases son: "

		for _, v := range c.lessons {
			txt += v + ", "
		}

		return txt[:len(txt)-2]
	}

	Go := Course{ // * Instancia de estructura
		"Curso Go",
		"Alex",
		19.00,
		false,
		[]uint{12, 33},
		map[uint]string{
			1: "Primeros pasos",
			2: "Conceptos basicos",
			3: "Programacion estructurada",
		},
	}

	Css := Course{
		name:  "Curso CSS",
		price: 24.00,
	}

	js := Course{}

	js.name = "Curso JavaScript"
	js.lessons = map[uint]string{
		1: "Introduccion a JS",
	}

	js.lessons[2] = "Propiedades JS"

	fmt.Printf("%+v", Go)
	fmt.Println(Go.lessons)

	fmt.Println(Css)
	fmt.Printf("%+v\n", js)
	GetClases(Go)

	(Go).ChangePrice(16)
	fmt.Println(Go)

	fmt.Println("////////")

	BBDD := new(Course)

	fmt.Printf("%+v", BBDD)

	BBDD.name = "Bases de datos"

	Otro := new(Course)

	fmt.Println(Otro)

	otherCourse := Course{}

	fmt.Printf("%+v\n", otherCourse)

	newCourse := Course{
		name: "Nuevo curso",
		free: false,
		lessons: map[uint]string{
			1: "Introduccion",
		},
	}

	fmt.Printf("%+v\n", newCourse)

	testCourse := Course{
		price: 0,
		id: []uint{
			1, 2, 3,
		},
		name: "Logica de programacion",
		free: true,
		lessons: map[uint]string{
			1: "Introduccion",
		},
		professor: "Joe",
	}

	fmt.Printf("%+v\n", testCourse)

	Html := Course{
		"HTML desde cero",
		"Mau González",
		19.99,
		false,
		[]uint{
			1, 2, 44, 12,
		},
		map[uint]string{
			1: "Introducción",
			2: "etiquetas",
		},
	}

	fmt.Printf("%+v", Html)
	clases := GetClases(Html)
	fmt.Println(clases)
	pokemon := games.Game{}
	fmt.Println(pokemon)
}

type Course struct { // * Estructura
	name, professor string
	price           float64
	free            bool
	id              []uint
	lessons         map[uint]string
}

func (c *Course) ChangePrice(p float64) {
	c.price = p
}
