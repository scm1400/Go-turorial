<br>

# Get Started With Go<br><br> From [노마드코더 아카데미](https://nomadcoders.co/go-for-beginners)

## 문법
---
### <b>1. 변수선언(Variables)<br><br>

```go
var name string = "muice" // 일반적 변수 선언 방법
myname := "muice" // 자동으로 타입을 지정해주는 변수 선언. func 안에서만 인식됨.
names := [5]string{"muice", "iceking", "ice"} // Array 변수 선언
```
<br>
`String name = "";`  ⚡ Java 에서 String 변수선언<br>
`String[5] id = new string[5]` ⚡ Java 에서 String Array 선언<br>
<br>

---
### <b>2. 반복문(Loop)<br><br>

- 일반적인 반복문 (...int는 배열을 의미)
```go
func superAdd(numbers ...int) int {

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
```
<br>
  
- array의 range로 반복문을 실행하는 경우  <br>
index를 작성하지않으면 (n-1)개의 값이 출력됨.<br>

```go
func superAdd(numbers ...int) int {

	for index, number := range numbers {
		fmt.Println(number)
	}
	return 1
}
```
-  `_` -> vlaue를 ignore 하는 방법 (변수가 선언되고 사용되지 않으면 코드실행이 안되기 때문에 사용)
```go
func superAdd(numbers ...int) int {
for _, number := range numbers {
	total += number
}
return total
}
```
<br>

### <b>3. if, switch문
```go
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

```

<br>

### <b>4. 배열(Array)

```go
// ⚡ 크기를 제한한 배열
names := [5]string{"muice", "iceking", "ice"}
names[4] = "new one"
// ⚡ 크기를 제한하지않은 배열(Slice라고 부름)
// append로 요소를 추가해줌
names2 := []string{"muice", "iceking", "ice"}
names2 = append(names2, "new one")
```

<br>

### <b>5. Map

```go
// ⚡ Key값이 [String], value가 String 타입인 Map 선언
muice := map[string]string{"name": "muice", "age": "12"}
fmt.Println(muice)
```

<br>

### <b>6. Struct

```go
// ⚡ person struct 생성
type person struct {
	name    string
	age     int
	favFood []string
}

// ⚡ person_muice에 값 할당
favFood := []string{"pizza", "chicken"}
person_muice := person{name: "muice", age: 18, favFood: favFood}
fmt.Println("person_struct : ", person_muice)
```

<br>

---
## <b>함수특징?
---

<br>
⚡ "naked" return : return할 값을 지정해주지 않아도 알아서 작동함<br>
⚡ defer : 함수가 값을 return 하고나서 작동함<br>

```go
func lenAndUpper(name string) (length int, uppercase string) {
    defer fmt.Println("i'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return //uppercase와 length를 리턴함
}

```
<br>
⭐ arguments로 여러값을 받아오려면 ...(3개)를 표시해야함

```go
func superAdd(numbers ...int) int {

}
```
