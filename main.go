package main

import (
	"GolangEd2024/arrays"
	"GolangEd2024/cycles"
	"GolangEd2024/numbers"
	"GolangEd2024/slice"
)

func main() {
	//Запуск функции из пакета numbers
	numbers.Numbers()
	//Запускаем функцию из пакета slice
	slice.Slice()
	//Запускаем функцию из пакета cycles
	cycles.Loops()
	//Запускаем функцию из пакета arrays
	arrays.Arrays()

}
