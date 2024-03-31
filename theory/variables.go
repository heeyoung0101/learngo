package theory

import "fmt"

// isMan := true 축약형은 func 함수 안 + 변수인 경우만 가능
var isMan bool = true

// private 함수
func sayBye() {
	fmt.Println("Bye")
}

// export 함수
func SayHello() {
	fmt.Println("hello")
}

func GetVariable() {
	const name string = "heeyoung" // 상수
	location := "America"          // var + 타입은 go가 찾아줌(var location string = "America")
	location = "Korea"             // 변수
	yesOrNo := true

	fmt.Println(name)
	fmt.Println(location)
	fmt.Println(yesOrNo)
}
