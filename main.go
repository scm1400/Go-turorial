package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

// ⚡ "naked" return : return할 값을 지정해주지 않아도 알아서 작동함
// ⚡ defer : 함수가 값을 return 하고나서 작동함
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("i'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// ⭐ arguments로 여러값을 받아오려면 ...(3개)를 표시해야함
func superAdd(numbers ...int) int {

	// 일반적인 Loop
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	//array의 range로 반복문을 실행하는 경우
	//index를 작성하지않으면 (n-1)개의 값이 출력됨.
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }

	total := 0
	// 💥 '_' -> vlaue를 ignore 하는 방법 (변수가 선언되고 사용되지 않으면 코드실행이 안되기 때문에 사용)
	for _, number := range numbers {
		total += number
	}
	return total

}

func canIDrink(age int) bool {

	// ⚡ if 조건에 변수선언을 할 수 있다.
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	// ⚡ switch 에서도 변수선언을 할 수 있다.
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	case 50:
		return true
	}
	return true

}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	totalLength, uppercase := lenAndUpper("muice")
	fmt.Println("lenAndUpper func : ", totalLength, uppercase)

	result := superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("superAdd func : ", result)

	fmt.Println("canIDrink func : ", canIDrink(16))

	// ⚡ 포인터
	a := 2
	b := &a
	*b = 20
	fmt.Println("Pointer : ", a, b)

	// ⚡ 크기를 제한한 배열
	names := [5]string{"muice", "iceking", "ice"}
	names[4] = "new one"
	// ⚡ 크기를 제한하지않은 배열(Slice라고 부름)
	// append로 요소를 추가해줌
	names2 := []string{"muice", "iceking", "ice"}
	names2 = append(names2, "new one")

	// ⚡ Map
	muice := map[string]string{"name": "muice", "age": "12"}
	fmt.Println(muice)

	// ⚡ Struct
	favFood := []string{"pizza", "chicken"}
	person_muice := person{name: "muice", age: 18, favFood: favFood}
	fmt.Println("person_struct : ", person_muice)

}
