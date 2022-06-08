package main

import "io/ioutil"

func main() {
	file, _ := ioutil.ReadFile("D:\\hello123.txt")

	_ = ioutil.WriteFile("D:\\hello.txt", file, 0666)

}
