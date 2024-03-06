package main

import "fmt"

func main() {
	d := make(map[string]string)
	d["Something"] = "Something"

	fmt.Printf("%#v\n", d)

	s := d["other"]
	fmt.Printf("%#v\n", s)
}
