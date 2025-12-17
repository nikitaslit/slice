package main

import (
	"fmt"
	"math/rand"
)

func slice(slice []int) { //копия
	//fmt.Println(slice)
}
func main() {

	var sres1 [4]int = [4]int{5, 1, 2, 5}

	var sres2 [6]int = [6]int{4, 2, 5, 1, 1, 2}

	var sres11 [6]int = [6]int{5, 1, 2, 5}

	var sres12 [6]int = [6]int{4, 2, 5, 1, 1, 2}

	var sres13 [4]int = [4]int{5, 1, 2, 5}

	var sres14 [6]int = [6]int{4, 2, 5, 1, 1, 2}

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

	//Найти пересечения значений в двух слайсах.
	slice111 := make([]int, 0, cap(sres11))
	//fmt.Println(slice111)
	//fmt.Println(slice1212, " ", slice2121)
	for i := 0; i < len(slice1212); i++ {
		for i1 := 0; i1 < len(slice2121); i1++ {
			if slice1212[i] == slice2121[i1] {
				slice111 = append(slice111, slice2121[i1])
				//fmt.Println(slice2121[i1], "нужен")
				slice2121 = append(slice2121[:i1], slice2121[i1+1:]...)
				slice1212 = append(slice1212[:i], slice1212[i+1:]...)
				fmt.Println(slice1212, " ", slice2121)
				i-- //если удалили нужно вернутсья

				//if i1 != 0 {
				//	i1--
				//}
				//if i != 0 {
				//	i--
				//}
				//fmt.Println("index", i, " ", i1)
				fmt.Println(slice111, "новый")
				break
			}
			//fmt.Println("index", i, " ", i1)
			//fmt.Println(slice1212)
		}
	}
	fmt.Println("Пересечение:", slice111)

	//Вернуть слайс с уникальными элементами из слайсов a и b.

	slice222 := append(sres13[:], sres14[:]...)
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
	//Если функция, то нужно передавать по ссылке
	//Если передавать по функции то слайс пересоздатся то оригинал тоже будет новый
	//Если будем работать с копией исходный слайс этого не затронет
	//Изначально слайс увеличивает свой Capasity 2x
	//Если Capasity больше 1024 то только на 25%
	slice(slice222)
}
