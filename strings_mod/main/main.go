//Downtiser
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	str := "Downtiser is a noob"
	new_str := strings.Replace(str, " ", "#", -1)
	fmt.Println(str, new_str)
	fmt.Println(strings.Count(str, "o"))
	fmt.Println(strings.Repeat("no", 5))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.Trim("nonoob is a noon", "no"))
	fmt.Println(strings.Fields(str))
	fmt.Println(strings.Split("a*b*c", "*"))
	fmt.Println(strings.Join([]string{"a", "b", "c"},  "*"))
	fmt.Println(strconv.Itoa(15))
	fmt.Println(strconv.Atoi("15"))
}