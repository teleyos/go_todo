package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func owo() {
	for _, item := range []interface{}{5,5.1,"mystring"} {
		fmt.Printf("%v %T\n",item,item)
	}
}

func help(code int){
	if code > 0 {
		fmt.Printf("error %d : ",code)
	}
	switch code {
	case 1:
		fmt.Printf("bitch, at least 1 arg pls")
	case 2:
		fmt.Printf("bitch, i don't do wtf '%v' is",os.Args[1])
	case 3:
		fmt.Printf("bitch, at least 2 arg pls")
	case 4:
		fmt.Printf("hum, actualy (spoiler alert) '%v' isn't a fkin number dumbass",os.Args[2])
	default:
		fmt.Printf("something is wrong (you)")
	}
	os.Exit(1)
}

func add(args []string){
	fmt.Printf("added '%v' to your todos",strings.Join(args," "))
}

func remove(id int){
	fmt.Printf("suppression de todos[%v]",id)
}

func main(){
	args := os.Args[1:]

	if len(args) < 1 {
		help(1)
	}

	switch args[0] {
	case "add":
		if len(args) < 2 { help(3) }
		add(args)
	case "remove":
		if len(args) < 2 { help(3) }
		num, err := strconv.Atoi(args[1])
		if err != nil { help(4) }
		remove(num)
	default:
		help(2)
	}
}

