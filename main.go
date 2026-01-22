package main

import "fmt"

func init() {
	fmt.Println("test ... 30/05/2002")
}
func main() {
	fmt.Println("hello , world")
	myAge := 10
	myLanguage := "Thai"
	const data string = "sawasdekab tansamachick chomrom kon chob hee"
	fmt.Println(myAge)
	fmt.Println(myLanguage)
	fmt.Println(data)

	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		} else if i == 10 {
			break
		}
		fmt.Println(i)
	}
}
