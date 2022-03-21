package linux

import (
	"flag"
	"fmt"
)

var (
	boolFlag   bool
	intFlag    int
	stringFlag string
)

func register() {
	flag.BoolVar(&boolFlag, "bool", false, "help message for \"b\" option")
	flag.BoolVar(&boolFlag, "b", false, "help message for \"b\" option")
	flag.IntVar(&intFlag, "int", 1234, "help message for \"i\" option (default 1234)")
	flag.IntVar(&intFlag, "i", 1234, "help message for \"i\" option (default 1234)")
	flag.StringVar(&stringFlag, "string", "defalut", "help message for \"s\" option (default \"default\")")
	flag.StringVar(&stringFlag, "s", "defalut", "help message for \"s\" option (default \"default\")")

	flag.Parse()
}

func Do() {
	fmt.Println("aaa")
	register()

	fmt.Println(boolFlag)
}
