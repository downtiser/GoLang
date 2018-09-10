//Downtiser
package add_package

import (
	"fmt"
	_ "go_dev/day2/example/test"
)

var Name string
var Age int

func init()  {
	Name = "Downtiser"
	Age = 20
	fmt.Println("The init add")
}