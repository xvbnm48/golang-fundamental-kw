package main

import "fmt"

type Student struct {
	ID    string
	Name  string
	Age   int
	Email string
}

func main() {
	student := Student{}
	student.ID = "123"
	student.Name = "sakura"
	student.Age = 20
	student.Email = "sakuraendo@gmail.com"
	fmt.Println(student)

	data := Student{"1", "sakura", 20, "sakuramiyawaki@gmail.com"}
	fmt.Println(data.Name)

	// sturct with parameter
	// data1 := displayStruct(data)
	// fmt.Println(data1)

	// method
	var s1 = Student{"1", "sakura", 20, "sakura@gmail.com"}
	s1.dataStudent()
	s1.getName()
}

func displayStruct(user Student) string {
	return fmt.Sprintf("name is %s ,with age %d and you email is %s", user.Name, user.Age, user.Email)

}

func (s Student) dataStudent() {
	fmt.Println("name is ", s.Name, "with age ", s.Age, "and you email is ", s.Email)
}

func (s Student) getName() {
	fmt.Println(s.Name)
}
