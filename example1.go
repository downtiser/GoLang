//Downtiser
package main

import(
	"fmt"
	add "go_dev/day2/example/add_package"
)

var Name string = "Noob"
var Age int = 15

func init()  {
	fmt.Println("The init main")
}

func output(n int)  {
	for i:= 0;i <= n ;i++  {
		fmt.Printf("%d + %d = %d\n",i, 5-i, n)
	}
}

func main()  {
	output(5)
	fmt.Printf("The add:Name:%s Age:%d\n", add.Name, add.Age)
	fmt.Printf("The main:Name:%s Age:%d",Name, Age)
}