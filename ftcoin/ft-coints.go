package main

import (
	"fmt"
	"math"
)

func Ft_coin(coins []int, amount int) int {
	// je un crèe un tableau pour stocker le nombre minimum de pièces pour chaque montant
	x := make([]int, amount+1)
	for i := range x {
		x[i] = math.MaxInt32
	}
	x[0] = 0

	// je parcous chaque montant jusqu'à amount
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 && x[i-coin] != math.MaxInt32 {
				x[i] = int(math.Min(float64(x[i]), float64(x[i-coin]+1)))
			}
		}
	}

	if x[amount] == math.MaxInt32 {
		return -1
	}
	return x[amount]
}

func main() {
	fmt.Println(Ft_coin([]int{1, 2, 5}, 11)) // resultat : 3
	fmt.Println(Ft_coin([]int{2}, 3))        // resultat : -1
	fmt.Println(Ft_coin([]int{1}, 0))        // resultat : 0
}
