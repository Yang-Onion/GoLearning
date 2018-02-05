package main

import "fmt"

func funcArr(){
	var arr [5]int
	fmt.Println("emp:",arr)

	arr[4]=100
	fmt.Println("set:",arr)
	fmt.Println("get:",arr[4])

	fmt.Println("length:",len(arr))


	b:=[5]int{1,2,3,4,5}
	fmt.Println("dcl:",b)

	var secArr [2][3]int

	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			secArr[i][j] =i+j
		}
	}
	fmt.Println("2d:",secArr)
}