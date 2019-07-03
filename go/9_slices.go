package main

import "fmt"

func main(){
	// var name make([]string , 6)			error
	name:= make([]string,6)
	fmt.Println("original string", name)

	// name[0] = 's'					error
	name[0], name[1] = "s", "u"

	// name[2:5] = "shil"			error
	fmt.Println("original string after assigning elements:",name)

	name = append(name, " bhi","l","e")
	fmt.Println("original string after append",name)

	name2:= make([]string, len(name))
	copy(name2, name)

	fmt.Println("copied string: ",name2)
	fmt.Println("printing 2:5",name[0:2])

	n1 := name[0:2]
	fmt.Println("name[0:2]:- ",n1)

	n1 = name[:len(name)]
	fmt.Println("name[:len(name)]:- ",n1)

	n1 = name[5:]
	fmt.Println("name[5:]:- ",n1)

	// oneD := make([]string{"sushil sanjay bhile"})		error because make is used to create only empty slices
	oneD := []string{"s","u","s","h","i","l"," ","s","a","n","j","a","y"," ","b","h","i","l","e"}
	fmt.Println(oneD)

	// twoD := make([][]int,5,3)		error	it takes last digit as size of outermost parent array. in this case it becomes [3][]

	twoD := make([][]int,5)
	for i:=0;i<5;i++{
		twoD[i] = make([]int, i)
		for j:=0; j<i; j++{
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}