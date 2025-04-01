package main

func main() {
	// 方式1
	const a int = 1

	// 方式2
	const b = "test"

	// 方式3
	const c, d = 2, "hello"

	// 方式4
	const e, f bool = true, false

	// 方式5
	const (
		h    byte = 3
		i         = "value"
		j, k      = "v", 4
		l, m      = 5, false
	)

	const (
		n = 6
	)
}
