package studentmanage

import "fmt"

type Student struct {
	Id        int8
	Name      string
	Age       int8
	ClassName string
}

func (student Student) study(subject string) {
	fmt.Printf("%s正在学习%s科目。", student.Name, subject)
}
