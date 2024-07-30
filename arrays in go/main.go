package main

import ("fmt")

func main(){
	fmt.Println("Arrays in go")

	var s[5] int;
	fmt.Println(s)

	t:=[5] int{4,6,8,10,12}
	for z:= range t {
		fmt.Println(z+1,t[z])
	}

	apples :=[5] int{2,4,6,8,100}
	fmt.Println((len(apples)))
}