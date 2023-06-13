package no

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenOrderNo() string {
	//default 4位
	return GenOrderNoWithLength(4)
}

func GenOrderNoWithLength(size int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	//纳秒级别时间戳
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < 4; i++ {
		//随机input 数值
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			return ""
		}
	}
	return sb.String()
}
