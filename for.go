package main

import (
	"flag"
	"fmt"
)

type A struct {
	ID   int
	Name string
}

func (r *A) ChangeName(name string) {
	r.Name = name
	fmt.Println(*r)
}

func main() {
	//c := A{55, ""}
	//fmt.Println(c)
	//r := &c

	//flag.Parse()
	//r.ChangeName(flag.Arg(0))

	for i := 0; i <= 5; i++ {
		flag.Parse()
		a := A{i, flag.Arg(i)}

		if a.ID <= 3 && a.Name != "" {
			fmt.Println(a)
		} else {
			fmt.Println("error")
		}

	}
}
