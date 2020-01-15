package core

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func str(u uint64) (string, uint64) {
	s := strconv.FormatUint(u, 2)
	return strings.Repeat("0", 64-len(s)) + s, u
}

func TestHello(t *testing.T) {
	t.Run("<<", func(t *testing.T) {
		fmt.Println(str(uint64(1)))
		fmt.Println(str(uint64(3 << 1))) // 将 3 的二进制数向左移动一位
		fmt.Println(str(uint64(1 << 2)))
		fmt.Println(str(uint64(1 << 3)))
		fmt.Println(str(uint64(1)))
		fmt.Println(str(uint64(2 >> 1))) // 将 2 的二进制数向右移动一位
		fmt.Println(str(uint64(6)))
		fmt.Println(str(uint64(1 << 63)))
		fmt.Println(str(uint64(1<<63 - 1)))
	})

	var a interface{}

	switch v := a.(type) {
	case int:
		fmt.Println(v == 1)
	case string:
		fmt.Println(v == "string")
	}

	t.Run("xxx", func(t *testing.T) {
		fmt.Printf("%T\n", 'A') // rune
		fmt.Printf("%#v\n", struct {
			s string
		}{
			s: "1",
		})
		fmt.Println([]byte("ABC")) //

		fmt.Println(strconv.FormatUint(256, 2))
	})
}

func TestHello2(t *testing.T) {
	t.Run("test hello2", func(t *testing.T) {
		i, j := 100, 1024
		p := &i // 生成指针 *p，并指向其操作数 i
		q := &j
		fmt.Println(*p, *q) // 通过指针读取
	})
}
