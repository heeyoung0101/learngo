package theory

import "fmt"

func GetPointer() {
	a := 1
	b := &a // a의 메모리 주소를 b에 저장
	a = 2
	fmt.Println(&a, b) // 0x14000110040 0x14000110040
	fmt.Println(a, *b) // 2 2

	*b = 20        // b의 값을 바꾸면 a에도 적용
	fmt.Println(a) // 20

	// &a : a의 주소
	// *b : b의 값
}
