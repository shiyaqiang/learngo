package main

import "fmt"

func main(){
	if 7%2 ==0{
		fmt.Println("7 is seven")
	}else{
		fmt.Println("7 is odd")
	}

	if 8%4==0{
		fmt.Println("8 is divisible by 4")
	}
	if num :=9;num<0{
		fmt.Println(num,"is negetave")
	}else if num<10{
		fmt.Println(num,"is negative")
	}else {
		fmt.Println(num,"has nultiple digits")
	}
}

