package main
import "fmt"

type student struct {
	name string
	rollNo int
}

func createStudent(name string, rollno int) student {
	s := student{name: name, rollNo: rollno}
	return s
}

func main(){
	p := fmt.Println
	p("This programs illustrates structures in go")

	s := student{"sushil", 1}
	p(s)
	p(student{"sush", 2})
	p(createStudent("su", 3))

	s1 := createStudent("s", 4)
	p(s1.name)
	p(s1.rollNo)

	s1.rollNo = 5
	p(s1.rollNo)
}
