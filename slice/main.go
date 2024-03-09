package main

import (
	"fmt"
)

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
	//Создание пустого map
	createMap := make(map[string]int)
	fmt.Println(createMap)
	createMap["K1"] = 12
	createMap["K2"] = 15
	fmt.Println(createMap)
	//Создание map c данными
	iMap := map[string]int{
		"K3": 20,
		"k4": 30,
	}
	fmt.Println(iMap)
	//Удаление элемента map функция delete
	delete(iMap, "K3")
	fmt.Println(iMap)
	//Перебор элементов map
	for key, value := range createMap {
		fmt.Println(key, value)
	}

}
