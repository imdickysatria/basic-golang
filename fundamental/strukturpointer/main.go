package main

import "fmt"

type Student struct{
	ID int
	Name string
	GPA float32
}

func main()  {

	student := Student{1, "Agung", 3.42}

	fmt.Println(student.Name)

	// graduate(&student)
	student.graduate()
	fmt.Println(student.Name)
	
}

// func graduate(student *Student)  {
// 	student.Name = student.Name + " S.Kom"
// 	fmt.Println(student.Name)
// }

func (student *Student) graduate()  {
	student.Name = student.Name + " S.Kom"
}
