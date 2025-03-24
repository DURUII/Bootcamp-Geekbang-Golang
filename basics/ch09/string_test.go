package ch9

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	// string 不是引用/指针类型，空值为 ""
	/*
		```python
		'李'.encode('utf-8').hex()
		bytes.fromhex('e69d8e').decode('utf-8')
		```
	*/
	// string 是 `只读` 的 byte 切片
	// 因此 len 返回 byte 数
	var s string = "\xe6\x9d\x9c\xe7\x9d\xbf"
	t.Log(s, len(s))

	chars := []rune(s)
	t.Logf("杜 unicode %x", chars[0])
	t.Logf("杜 utf-8 %x", s[:3])
	// string 的 byte 数组可以存放任何数据
}

func TestString2(t *testing.T) {
	str := "hello"
	// cannot assign to str[0] (value of type byte)
	str[0] = 'x'
	fmt.Println(str)
}

func TestRune(t *testing.T) {
	s := "hello, 中国🀄️!"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
}
