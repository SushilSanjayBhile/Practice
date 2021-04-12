package main
import "fmt"

func main(){
	s := make([]string, 2)
	s[0] = "sush"
	s[1] = "il"

	append(s, " bhile")
	fmt.Println(s)
}
