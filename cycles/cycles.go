package cycles

import "fmt"

func Loops() {
	//Цикл for
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	//Использования ключевого слова range
	var anArray = [5]int{0, 1, -1, 2, -2}
	for index, value := range anArray {
		fmt.Println("index:", index, "value:", value)
	}
}
