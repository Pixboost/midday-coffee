package main

import (
	"fmt"
	"html/template"
	"os"
)

type CoffeeCup struct {
	Name string
	Price float32
	Image string
}

func main()  {
	tmpl, err := template.New("cups.tmpl.html").ParseFiles("web/original/cups.tmpl.html")
	if err != nil {
		fmt.Printf("couldn't parse template: %v\n", err)
		os.Exit(1)
	}

	f, err :=os.OpenFile("web/original/cups.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("couldn't create output file: %v\n", err)
		os.Exit(2)
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		fmt.Printf("couldn't execute template: %v\n", err)
		os.Exit(3)
	}
}

var data = struct {
	Cups []CoffeeCup
}{
	Cups: []CoffeeCup{
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/blur-2846257__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/cappuccino-593256__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/cappuccino-756490__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/code-geek-2680204__340.png",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/coffee-386878__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/coffee-777612__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/coffee-2127225__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/coffee-3129995__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/coffee-3139776__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/coffee-3183337__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/coffee-3183729__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/coffee-3219749__340.jpg",
		},

		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/coffee-beans-3093228__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/coffee-mugs-459324__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/cup-1891446__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/cup-2315565__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/cup-2318315__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/cup-3082145__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/cup-3117730__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/cup-3183995__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/keep-calm-2816357__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/mug-3059994__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/smile-2001662__340.jpg",
		},
		{
			Name:  "White",
			Price: 9.99,
			Image: "../assets/cups/tea-cup-2107599__340.jpg",
		},
	},
}

