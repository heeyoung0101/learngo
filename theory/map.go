package theory

import "fmt"

func GetMap() {
	heeyoung := map[string]string{"name": "heeyoung", "age": "29"}
	// key와 value 모두 초기화한대로 string만, "age": 29 는 불가능 -> struct에서 가능
	fmt.Println(heeyoung)

	for key, _ := range heeyoung { // value 숨긴 채로 key만 반복
		fmt.Println(key)
		// name
		// age
	}
}
