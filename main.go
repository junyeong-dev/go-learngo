package main

import (
	"fmt"

	"github.com/JunYeong-dev/learngo/something"
)

// Node.js, python과 다르게 특정 function을 찾게 되는데, 그게 바로 function main - func main()
// 컴파일러가 main package와 main function을 먼저 찾고 실행시킴
func main() {
	// Println과 같이 왜 함수의 시작이 대문자일까?
	fmt.Println("Hello World")
	// 이처럼 시작이 대문자인 것들만 package로 사용이 가능
	something.SayHello()

	// 상수(Constant)와 변수(Variables)
	// Go는 type언어 이기 때문에 type이 무엇인지 설정해 줘야 함(Java나 C언어와 같이)
	const name string = "nick"
	// 상수이기 때문에 변경 불가능
	// name = "judy" (x)

	// 변수기 때문에 변경 가능
	var a string = "nick"
	// a := "nick" <- 축약형; 이 경우에는 Go가 알아서 type을 찾아줌
	// 축약형은 오로지 function안에서만 가능하고, 변수에만 적용 가능
	a = "judy"
	fmt.Println(a)
}
