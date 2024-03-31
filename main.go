package main // package 명시 필수

import (
	"fmt" // vscode는 자동 import, fmt 포맷팅을 위한 패키지
	"strings"

	"github.com/heeyoung0101/learngo/something"
)

// isMan := true 축약형은 func 함수 안 + 변수인 경우만 가능
var isMan bool = true

func multiply(a, b int) int { // 파라미터와 결괏값의 타입 정하기
	return a * b
}

func lenAndUpper(name string) (int, string) { // go는 여러개의 결과를 리턴할 수 있음
	return len(name), strings.ToUpper(name)
}
func repeatMe(words ...string) { // 타입 앞에 ...을 찍어 원하는 만큼 arguments 받기
	fmt.Println(words)
}

func lenAndUpper2(name string) (length int, uppercase string) { // 두 개의 변수 미리 만들기
	defer fmt.Println("I'm done") // function의 retern하고 실행시킬 코드(콜백함수 같은듯?)
	length = len(name)
	uppercase = strings.ToUpper(name)
	return // 두 개의 변수 리턴(명시하지 않아도 됨)
}

func main() { // go 프로그램의 시작점(필수!!)
	fmt.Println("Hello World!")

	something.SayHello() // export된 메서드는 대문자로 시작
	// something.sayBye() private 메시드여서 사용할 수 없음

	const name string = "heeyoung" // 상수
	location := "America"          // var + 타입은 go가 찾아줌(var location string = "America")
	location = "Korea"             // 변수
	yesOrNo := true
	fmt.Println(name)
	fmt.Println(location)
	fmt.Println(yesOrNo)

	fmt.Println(multiply(2, 5))

	totalLenth, upperName := lenAndUpper("heeyoung") // 두 개의 결과를 두 변수에 나눠서 받기
	totalLenthOnly, _ := lenAndUpper("heeyoung")     // 두 번째 변수는 무시(컴파일되지 않음)
	fmt.Println(totalLenth, upperName)               // 8 HEEYOUNG
	fmt.Println(totalLenthOnly)

	repeatMe("heeyoung", "gwidong") // [heeyoung gwidong]

	totalLength, up := lenAndUpper2("gwidong")
	fmt.Println(totalLength, up) // 7 GWIDONG
}