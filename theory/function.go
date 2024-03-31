package theory

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int { // 파라미터와 결괏값의 타입 정하기
	return a * b
}

// go는 여러개의 결과를 리턴할 수 있음
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 타입 앞에 ...을 찍어 원하는 만큼 arguments 받기
func repeatMe(words ...string) {
	fmt.Println(words)
}

// 두 개의 결괏값 변수 미리 만들기
func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // function의 retern하고 실행시킬 코드(콜백함수 같은듯?)
	length = len(name)
	uppercase = strings.ToUpper(name)
	return // 두 개의 변수 리턴(명시하지 않아도 됨)
}

func GetFunction() {

	fmt.Println(multiply(2, 5))

	totalLenth, upperName := lenAndUpper("heeyoung") // 두 개의 결과를 두 변수에 나눠서 받기
	totalLenthOnly, _ := lenAndUpper("heeyoung")     // 두 번째 변수는 무시(컴파일되지 않음)
	fmt.Println(totalLenth, upperName)               // 8 HEEYOUNG
	fmt.Println(totalLenthOnly)

	repeatMe("heeyoung", "gwidong") // [heeyoung gwidong]

	totalLength, up := lenAndUpper2("gwidong")
	fmt.Println(totalLength, up) // 7 GWIDONG
}
