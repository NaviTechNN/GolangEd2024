package main

import "fmt"

func main() {
	//Создание среза(slice, mutablelist)
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	//Вывод на экран среза
	fmt.Println(a6)
	fmt.Println(a4)
	//Создание пустого среза
	var asSlice = make([]int, 5)
	//Вывод на экран пустого среза
	fmt.Println(asSlice)
	//Добавление элементов в срез функция append
	asSlice = append(asSlice, 5)
	fmt.Println(asSlice)

}
