package Iteration

const repeatCount = 5

func Repeat(s string) string {
	res := ""
	for i := 0; i < repeatCount; i++ {
		res += s
	}
	return res
}
