package main

import "fmt"

type Hello struct {
	hello string
}

func main() {
	fmt.Println("Hello from main")
	helloFromFunc()
	fmt.Println(helloReturningFromFunc)
	fmt.Println(helloReturningFromFunc())
	helloSentToFunc("Hello sent to func")

	var greet = func(s string) {
		fmt.Println("Hello", s)
	}

	greet("from variable")

	x := helloFromReturningFunc("Hello")

	fmt.Println(x("\tfrom duck"))
	fmt.Println(x("\tfrom cow"))

	fmt.Println(helloFromParamFunc("Hello", "from", dash))
	fmt.Println(helloFromParamFunc("Hello", "from", colon))

	func() {
		fmt.Println("Hello from anonymous func")
	}()

	helloFromManyFunc("many", "many", "many", "many")

	c := closuresFunc()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

	fmt.Println(recursiveFunc(15, "Hello"))

	h := Hello{hello: "Hello"}
	h.helloFromStruct()

	fmt.Println(implicitReturnFunc())
}

func implicitReturnFunc() (h1, h2 string) {
	h1 = "Hello"
	h2 = "\tfrom implicit func"
	return
}

func (h Hello) helloFromStruct() {
	fmt.Println(h.hello, "\tthere")
}

func recursiveFunc(i int, s string) string {
	if i == 0 {
		return s
	} else {
		return s + recursiveFunc(i-1, "\t"+s)
	}
}

func closuresFunc() func() string {
	s := "Hello"
	return func() string {
		s += "\tagain"
		return s
	}
}

func helloFromManyFunc(s ...string) {
	s1 := "Hello"
	for _, x := range s {
		s1 += "\t" + x
	}

	fmt.Println(s1)
}

func helloFromParamFunc(s1, s2 string, o func(string, string) string) string {
	return o(s1, s2)
}

func dash(s1, s2 string) string {
	return s1 + "\t-\t" + s2
}

func colon(s1, s2 string) string {
	return s1 + "\t:\t" + s2
}

func helloFromFunc() {
	fmt.Println("Hello from func")
}

func helloReturningFromFunc() string {
	return "Hello from returning func"
}

func helloSentToFunc(s string) {
	fmt.Println(s)
}

func helloFromReturningFunc(s string) func(string) string {
	return func(t string) string {
		return s + t
	}
}
