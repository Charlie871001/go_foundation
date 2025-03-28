package main

import "fmt"

func main() {
	// // 十六进制
	// var a uint8 = 0xF
	// var b uint8 = 0xf

	// // 八进制
	// var c uint8 = 017
	// var d uint8 = 0o17
	// var e uint8 = 0o17

	// // 二进制
	// var f uint8 = 0b1111
	// var g uint8 = 0b1111

	// // 十进制
	// var h uint8 = 15

	// var float1 float32 = 10.8
	// var float2 float64 = 10.09

	// var c1 complex64
	// c1 = 1.10 + 0.1i
	// c2 := 1.10 + 0.1i
	// c3 := complex(1.10, 0.1) // c2与c3是等价的

	// byte 是 uint8 的内置别名，可以把 byte 和 uint8 视为同一种类型。
	// 在 Go 中，字符串可以直接被转换成 []byte（byte 切片）。
	var s string = "Hello, world!"
	var bytes []byte = []byte(s)
	fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)
	fmt.Println("string length:", len(s))
	fmt.Println("bytes length:", len(bytes))

	// 同时[]byte 也可以直接转换成 string。
	var bytes1 []byte = []byte{72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}
	var s1 string = string(bytes1)
	fmt.Println(s1)

	// rune 是 int32 的内置别名，可以把 rune 和 int32 视为同一种类型。但 rune 是特殊的整数类型。
	// 一个 rune 类型的值由一个个被单引号包住的字符组成，比如：
	// var a rune = 'a'
	// var b rune = '是'

	// 字符串可以直接转换成 []rune（rune 切片）
	var s2 string = "Hello, world!"
	var runes []rune = []rune(s2)
	fmt.Println("convert \"Hello, world!\" to runes: ", runes)
	fmt.Println("runes length:", len(runes))

	// 在 Go 中，字符串是 UTF-8 编码的，并且所有的 Go 源码都必须是 UTF-8 编码。
	var s3 string = "Hello\nworld!\n"
	var s4 string = `Hello
world!
`
	fmt.Println(s3 == s4)

	var s0 string = "Go语言"
	var bytes0 []byte = []byte(s0)
	var runes0 []rune = []rune(s0)

	fmt.Println("string length: ", len(s0))
	fmt.Println("bytes length: ", len(bytes0), bytes0)
	fmt.Println("runes length: ", len(runes0), runes0)

	var s5 string = "Go语言"
	var bytes5 []byte = []byte(s5)
	var runes5 []rune = []rune(s5)

	fmt.Println("string sub: ", s5[0:7])
	fmt.Println("bytes sub: ", string(bytes5[0:7]))
	fmt.Println("runes sub: ", string(runes5[0:3]))

}
