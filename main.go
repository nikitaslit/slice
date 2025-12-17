package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var sres1 [4]int = [4]int{5, 1, 2, 5}

	var sres2 [6]int = [6]int{4, 2, 5, 1, 1, 2}

	var sres11 [6]int = [6]int{}

	var sres12 [6]int = [6]int{}

	//рандомно заполнить
	for i := 0; i < cap(sres1); i++ {
		sres11[i] = rand.Intn(100)
	}
	for i := 0; i < cap(sres2); i++ {
		sres12[i] = rand.Intn(100)
	}
	fmt.Println(sres11, " ", sres12)

	//Вернуть два слайса с уникальными элементами.
	//slice1212 := make([]int, len(sres1), len(sres1))
	slice1212 := append(sres1[:])
	for i := 0; i < len(slice1212); i++ {
		for i1 := 0; i1 < len(slice1212); i1++ {
			if (slice1212[i] == slice1212[i1]) && (i != i1) {
				slice1212 = append(slice1212[:i1], slice1212[i1+1:]...)
				i1-- //если удалили нужно вернутсья
				//fmt.Println(slice1212)
				break
			}
			//fmt.Println(slice1212)
		}
	}
	//slice2121 := make([]int, len(sres2), len(sres2))
	slice2121 := append(sres2[:])
	for i := 0; i < len(slice2121); i++ {
		for i1 := 0; i1 < len(slice2121); i1++ {
			if (slice2121[i] == slice2121[i1]) && (i != i1) {
				slice2121 = append(slice2121[:i1], slice2121[i1+1:]...)
				i1-- //если удалили нужно вернутсья
				//fmt.Println(slice2121)
				break
			}
			//fmt.Println(slice2121)
		}
	}
	fmt.Println(slice1212, " ", slice2121)

	slice111 := make([]int, cap(sres12), cap(sres12))
	//Найти пересечения значений в двух слайсах.
	for i := 0; i < len(sres1); i++ {
		for i1 := 0; i1 < len(sres2); i1++ {
			if sres1[i] == sres2[i1] {
				slice111[i1] = sres1[i]
			}
		}
	}
	fmt.Println(slice111)

	//Вернуть слайс с уникальными элементами из слайсов a и b.
	slice222 := append(sres1[:], sres2[:]...)
	fmt.Println(slice222)
	for i := 0; i < len(slice222); i++ {
		for i1 := 0; i1 < len(slice222); i1++ {
			if (slice222[i] == slice222[i1]) && (i != i1) {
				slice222 = append(slice222[:i1], slice222[i1+1:]...)
				i1-- //если удалили нужно вернутсья
				i--  //лишнее
				break
			}
		}
	}
	fmt.Println(slice222)
}
