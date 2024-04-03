package theory

import "fmt"

func GetArray() {
	names := [5]string{"nico", "lynn", "dal"} // [...]는 초기화 시 배열 길이를 자동 셋팅
	names[3] = "alala"
	names[4] = "alala"
	// names[5] = "alala"  범위를 벗어남
	fmt.Println(names)
}

func GetSlice() {
	names := []string{"nico", "lynn", "dal"} // []안을 비워둠
	names = append(names, "heeyoung")        // append는 새로운 값이 추가된 slice를 리턴
	fmt.Println(names)
}
