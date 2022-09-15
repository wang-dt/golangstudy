package main

var temp int

/**
 *wangdt
 */
//字符串交换
func changeval(a *int, b *int) {

	temp = *a
	*a = *b
	*b = temp
}

func main() {

	var a int = 10
	var b int = 20

	changeval(&a, &b)

	println(a, b)
}
