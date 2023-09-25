package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time // TIme 메서드를 호출한다?
	now = time.Now()
	//var year int = now.Year()
	var year = now.Year()
	month := now.Month()
	fmt.Println(year, month, now.Day(), now.Hour().now.Minute(), now.Second())
}

// 그러면 golang 이라는 파일 안에 go mod init은 한번만 실행 할 수 있나?
