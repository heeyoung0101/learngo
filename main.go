package main // package 명시 필수

import (
	"fmt" // vscode는 자동 import, fmt 포맷팅을 위한 패키지

	"github.com/heeyoung0101/learngo/theory"
)

func main() { // go 프로그램의 시작점(필수!!)
	fmt.Println("Hello World!")

	// #1. theory
	theory.SayHello() // export된 메서드는 대문자로 시작
	// theory.sayBye() private 메시드여서 사용할 수 없음

	// variables
	theory.GetVariable()

	// function
	theory.GetFunction()

	// for
	theory.GetFor()

	// if, switch
	theory.GetIf()
	theory.GetSwitch()

	// pointer
	theory.GetPointer()

	// array, slice
	theory.GetArray()
	theory.GetSlice()

	// map
	theory.GetMap()
}
