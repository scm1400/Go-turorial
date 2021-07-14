package banking

// ⭐account의 앞글자를 대문자로 해주어야 Public 인식한다.
// 소문자이면 private

type Account struct {
	Owner   string
	Balance int
}
