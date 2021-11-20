package main

import "fmt"

type alcoholInML int

type person struct{
	name string
	age int
	heightInCm float64
	bloodGroup string
}

type secretAgent struct{
	person 
	liscenceToKill bool
}

type human interface{
	speak()
}

func main() {
	x:= 7
	fmt.Printf("%T",x) // %T => inside of fmt.Printf gives the the "type"
	// Composite literals
	sliceOfInt := []int{2,4,6,8}
	fmt.Println(sliceOfInt)
	
	// map
	m := map[string]int{
		"kobe":24,
		"mj":23,
		"kd":35,
		"lbj":6,
	}
	fmt.Println(m)

	himanshu := person{"Himanshu Bisht",22,166.8,"AB+ve"}
	fmt.Println(himanshu)	
	himanshu.speak()
	// reciever function or methods 
	
	sherlockHolmes := secretAgent{person{"Sherlock Holmes", 42,180.3,"AB-ve",},true}
	sherlockHolmes.speak()  // this will return secretAgent type wala method
	sherlockHolmes.person.speak() // this will return person wala method

	// interface polymorphism
	saySomething(himanshu)
	saySomething(sherlockHolmes)

	// we can define our own types
	var chonti alcoholInML
	chonti = 180
	fmt.Println(chonti)
	fmt.Printf("%T",chonti)
}
func (p person) speak()  {
	fmt.Println(p.name + " is saying bad tings to you")
}

func (sa secretAgent) speak() {
	fmt.Println(sa.name + " is a cool guy")
}

func saySomething(h human){
	h.speak()
}