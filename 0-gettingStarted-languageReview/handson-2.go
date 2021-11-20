package main

import "fmt"

func main(){
	s:=`“i'm sorry dave i can't do that” to “s”`
	fmt.Println(s)
	bs := []byte(s)
	fmt.Println(bs)
	
}