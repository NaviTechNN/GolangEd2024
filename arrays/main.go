package main

import (
	"fmt"
)

func main() {
	//Массивы
	//Объявление массива
	var anArray = [4]int{0, 1, 2, 3}
	//Вывод длины массива
	fmt.Println(len(anArray))
	//Инициализация двухмерного массива
	var twoArray = [4][4]int{
		{0, 1, 2, 3},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	//Вывод длины двумерного массива
	fmt.Println(len(twoArray))
	//Сумма значений двумерного массива
	var sum int = 0
	//Перебор двумерного массива
	for i := 0; i < len(twoArray); i++ {
		for j := 0; j < len(twoArray[i]); j++ {
			fmt.Println("value: ", twoArray[i][j])
			sum = sum + twoArray[i][j]
		}

	}
	fmt.Println(sum)
	//Использование ключевого слова range
	for _, value := range twoArray {
		for _, twoValue := range value {
			fmt.Println("value:", twoValue)
		}
	}
}
