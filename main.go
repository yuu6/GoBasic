package main

import (
	"fmt"
	"os"
)

func main(){
	//println("hello world!")

	if len(os.Args) != 2{
		os.Exit(1)
	}
	// Go 的第一函数默认的是当前的执行路径

	age := 123

	age = 12

	fmt.Println("the age is",age);

	// struct

	//type People struct{
	//	Name string
	//	Power int
	//}
	//
	//yuhongchao:= People{"yuhongchao",12}
	//
	//fmt.Print(yuhongchao.Power)

	yuhongchao := &Saiyan{"yuhongchao",12}
	yuhongchao.Super()
	fmt.Print(yuhongchao.Power)


	// arrays
	var score [10]int
	score[0] = 12

	//scores := []int{2,23,4,5}

	//fmt.Print(scores[0])

	//testArray := make([]int, 0, 10)
	//scores = testArray[0:8]
	//fmt.Print(testArray)

	scores := make([]int , 0, 10)
	c := cap(scores)
	fmt.Println(c)

	for i := 0; i < 25 ; i++ {
		scores = append(scores, i)

		if cap(scores) != c {
			c = cap(scores)
			fmt.Println(c)
		}
	}



}

