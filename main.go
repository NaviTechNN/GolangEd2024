package main

import (
	"GolangEd2024/arrays"
	"GolangEd2024/cycles"
	"GolangEd2024/numbers"
	"GolangEd2024/slice"
	_struct "GolangEd2024/struct"
	"fmt"
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
	//Инициализация структуры Person
	p := _struct.Person{
		Name:   "Maksim",
		Height: 184,
		Weight: 92,
	}
	//Вывод структуры
	fmt.Println(p)
	//Вывод имени объекта Person
	fmt.Println(p.Name)
	//Создание структуры с помощью функции
	p1 := _struct.CreatePerson("Sergey", 200, 100)
	fmt.Println(p1)

}
