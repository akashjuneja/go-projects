package main
import (
	"fmt"
)

func main(){
	fmt.Println("Hello world")
	i:=1
	var x int =100

	var mango string="mango"
	fmt.Printf(mango,"/n")
	
	for i<=10 {
		fmt.Println(i)
		i++
	}

	for i:=5 ;i<=x;i=i+5{
		if i%2==0{
			fmt.Println(i)
		}
	
	}
	 
}