package theory

import "fmt"

// if 조건문
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 20 { // if문 안에서만 사용하기 위해 변수를 생성할 수 있음
		return false
	}
	return true
}

// switch 조건문
func canIDrink2(age int) bool {
	switch koreanAge := age + 2; koreanAge { // if-else 구문이 난무하는 것을 막아줌
	case 10:
		return false
	case 20:
		return true
	}
	return false
}

func GetIf() {
	fmt.Println(canIDrink(16))
}

func GetSwitch() {
	fmt.Println(canIDrink2(18))
}
