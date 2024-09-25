package main

import "fmt"

func Ft_missing(nums []int) int {
	longueur := len(nums)

	// je Calcule la somme théorique des nombres de 0 à longueur
	sommeTheorique := longueur * (longueur + 1) / 2
	// je Calcule la somme réelle des éléments présents dans nums
	sommeReelle := 0
	for _, nombre := range nums {
		sommeReelle += nombre
	}

	return sommeTheorique - sommeReelle
}

func main() {
	fmt.Println(Ft_missing([]int{3, 1, 2}))                   // Résultat attendu : 0
	fmt.Println(Ft_missing([]int{0, 1}))                      // Résultat attendu : 2
	fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // Résultat attendu : 8
}
