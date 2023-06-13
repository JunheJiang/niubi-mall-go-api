package str

import "fmt"

func SubStrLen(str string, length int) string {
	//byte uint8别名、 rune int32别名
	nameRune := []rune(str)
	//string 函数 len 函数 4字节
	fmt.Println("string(nameRune[:4]) = ", string(nameRune[:4]))
	if len(str) > length {
		return string(nameRune[:length-1]) + "..."
	}
	return string(nameRune)
}
