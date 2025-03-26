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

	var s string = "Hello, world!"
	var bytes []byte = []byte(s)
	fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)

	var bytes1 []byte = []byte{72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}
	var s1 string = string(bytes1)
	fmt.Println(s1)
}
