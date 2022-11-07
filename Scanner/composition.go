package main

import "fmt"

type details struct {
	genre        string
	genreRatring string
	reviews      string
}
type game struct {
	name  string
	price string
	details
}

func (d details) showDetails() {
	fmt.Println("Genre :", d.genre)
	fmt.Println("Genre Rating :", d.genreRatring)
	fmt.Println("Reviews : ", d.reviews)
}
func (g game) show() {
	fmt.Println("Name: ", g.name)
	fmt.Println("Price : ", g.price)
	g.showDetails()
}
