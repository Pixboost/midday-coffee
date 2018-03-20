package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"io/ioutil"
)

type CoffeeCup struct {
	Name string
	Price float32
	Image string
}

//Copies web/ folder to dist/ and builds all HTML templates.
//Template are identified by having *.tmpl.html suffix in the filename.
func main()  {
	err := os.RemoveAll("dist")
	if err != nil {
		fmt.Printf("couldn't remove dist dir: %v\n", err)
		os.Exit(4)
	}
	err = os.Mkdir("dist", os.ModeDir)
	if err != nil {
		fmt.Printf("couldn't create dist dir: %v\n", err)
		os.Exit(5)
	}

	filepath.Walk("web", func(path string, info os.FileInfo, err error) error {
		fmt.Printf("%s\n", path)
		distPath := fmt.Sprintf("dist/%s", path)
		if info.IsDir() {
			err = os.Mkdir(distPath, os.ModeDir)
			if err != nil {
				fmt.Printf("Couldn't create dist/%s dir: %v", path, err)
				os.Exit(6)
			}
		} else if strings.HasSuffix(path, "tmpl.html") {
			if err != nil {
				fmt.Printf("Couldn't read file %s dir: %v", path, err)
				os.Exit(7)
			}

			distFilename := strings.Replace(distPath, ".tmpl", "", 1)
			err = buildTemplate(path, filepath.Base(path), distFilename)
			if err != nil {
				fmt.Printf("Couldn't process template %s dir: %v", path, err)
				os.Exit(8)
			}
		} else {
			fileContent, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Printf("Couldn't read file %s dir: %v", path, err)
				os.Exit(7)
			}

			ioutil.WriteFile(distPath, fileContent, 0666)
		}
		return nil
	})
}

func buildTemplate(templateFilePath string, templateName string, outFilePath string) error {

	tmpl, err := template.New(templateName).ParseFiles(templateFilePath)
	if err != nil {
		fmt.Printf("couldn't parse template: %v\n", err)
		os.Exit(1)
	}

	f, err :=os.OpenFile(outFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
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

	return nil
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

