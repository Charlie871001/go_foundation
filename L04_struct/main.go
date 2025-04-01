package main

import (
	"fmt"
)

type A struct {
	a string
}

func (a A) string() string {
	return a.a
}

func (a A) stringA() string {
	return a.a
}

func (a A) setA(v string) {
	a.a = v
}

func (a *A) stringPA() string {
	return a.a
}

func (a *A) setPA(v string) {
	a.a = v
}

type B struct {
	A
	b string
}

func (b B) string() string {
	return b.b
}

func (b B) stringB() string {
	return b.b
}

type C struct {
	B
	a string
	b string
	c string
	d []byte
}

func (c C) string() string {
	return c.c
}

func (c C) modityD() {
	c.d[2] = 3
}

func callStructMethod() {
	var a A
	a = A{
		a: "a",
	}
	a.string()
}

func NewC() C {
	return C{
		B: B{
			A: A{
				a: "ba",
			},
			b: "b",
		},
		a: "ca",
		b: "cb",
		c: "c",
		d: []byte{1, 2, 3},
	}
}

func main() {
	c := NewC()
	cp := &c
	fmt.Println(c.string())
	fmt.Println(c.stringA())
	fmt.Println(c.stringB())

	fmt.Println(cp.string())
	fmt.Println(cp.stringA())
	fmt.Println(cp.stringB())

	c.setA("1a")
	fmt.Println("------------------c.setA")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)

	cp.setA("2a")
	fmt.Println("------------------cp.setA")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)

	c.setPA("3a")
	fmt.Println("------------------c.setPA")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)

	cp.setPA("4a")
	fmt.Println("------------------cp.setPA")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)

	cp.modityD()
	fmt.Println("------------------cp.modityD")
	fmt.Println(cp.d)
}

// 方法的接受者是值和指针的区别
// 声明结构体方法时，func 关键字后面以及方法名前面的括号中声明的这个类似变量的东西就是接受者，这个变量的类型是值就叫做值接受者，类型是指针就叫指针接受者。

// 不管接受者是什么类型都不影响变量调用方法。但是他们调用方法后，产生的结果会有一些不同。

// 值接受者（Value Receiver）: 值接受者的方法操作的是值的副本，在这个方法中可以随意修改值接受者的字段的值，但不会影响原始实例。

// 指针接受者（Pointer Receiver）：指针接受者的方法操作的是原实例的指针，修改指针接受者的任意字段，也意味着修改了原实例。

// 在实际代码中应该如何选择呢？

// 这个需要根据具体情况来定。

// 当结构体比较复杂并且数据量比较大时（比如持有一个[]byte，长度有 8000 以上），尽量使用指针接受者。因为执行值接受者方法时，会创建一个实例的副本，会占用比较大的栈空间，即使栈空间的清理很容易，但在一些资源受限的平台上运行，可能会导致内存爆掉，而被系统强制杀死进程。

// 当结构体只是存储数据，并且执行方法后，不希望方法中的临时修改影响到原来的实例，那么就可以选择使用值接受者
