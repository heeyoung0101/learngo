package theory

import "fmt"

// for 반복문
func superAdd(numbers ...int) int {
	total := 0
	// 같은 방식 for i:= 0; i < len(numbers); i++
	for _, number := range numbers { // index와 number, range는 for안에서만 작동
		fmt.Println(number)
		total += number
	}
	return total
}

func GetFor() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
