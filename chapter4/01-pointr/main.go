package main

import "fmt"

func main() {
	a, b := 1, 2
	add(&a, &b)
	fmt.Println(a)
	c := &a

	d := &c
	fmt.Println(d, *d, **d, &c)

	m := map[string]string{}
	mp1 := &m
	fmt.Println(mp1)
	put(m)
	f1 := add
	f1(&a, &b)
	f1p1 := &f1
	(*f1p1)(&a, &b)
	fmt.Println("f1p1,add =", a)

	//var noting *int
	//*noting = 3
	//fmt.Println(noting)

	{
		var notingMap map[string]string
		notingMap["a"] = "BBB"
		fmt.Println(notingMap)
	}

}
func put(m map[string]string) {
	m["a"] = "I am a"
}

func add(a, b *int) {
	*a = *a + *b
}
