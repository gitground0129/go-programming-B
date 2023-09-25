package main

import (
	"fmt" //이게 정리해주는 그건가..? 
	"math"
	"reflect"
	"strings"
)

func main() {
	var c7 string //
	var d int
	var e bool
	var f float64
	var G = 99

	fmt.Println(c7, d, e, f, G)

	//var a int // 선언
	//a = 7     // assign value

	//var a int = 7 // declaration & assign value

	//var a = 7 // declaration & assign value 한번 쓰인 타입은 쭉 쓰인다. int 를 선어하지 않아도 오른쪽 값으로 인식한다.
	a := 7
	fmt.Println(a, reflect.TypeOf(a))

	//b := 8.34 //float 64 : 64면 8바이트.
	var b float32 = 8.34
	fmt.Println(a * int(b))
	fmt.Println(float32(a) > b)

	fmt.Println(b, reflect.TypeOf(b)) // 내가 원룸 살고 싶은데 투룸은 너무 커

	fmt.Println('Z', '2', '\n', '김', '인', '하')   // 룬은 단일 문자를 가진다. (형)                                                                          //run literals(int32) 90, 50, 10, 44608, 51064, 54616
	fmt.Println(reflect.TypeOf('Z'), reflect.TypeOf(2), reflect.TypeOf("Hi"), reflect.TypeOf(4.99), reflect.TypeOf(false)) // doble 이란 개념 x float으로 처리
	fmt.Println(math.Floor(2.17), math.Ceil(2.17), math.Sqrt(16))                                                          //    \t 는 탭을 의마하고 \" \" 큰 따옴표를 활용한다.
	fmt.Println(strings.Title("open source\t programming\n \"go\"!"))                                                      //title은 첫글자를 대문자로 변환해준다.
	//fmt.Println(strings.Title("오픈소스 프로그래밍"))
	//fmt.Println("OpenSource Programming~")

	//koreanzombie := "정찬성" // 문법적으로 문제는 없지만  camel케이스 대로 사용하자
	koreanZombie := "정찬성"
	fmt.Println(koreanZombie) }