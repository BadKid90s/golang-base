package studentmanage

import "fmt"

type Manage struct {
	StudentList     []Student
	CourseList      []Course
	StudentScoreMap map[int]Score

	menuNum *int
}

func NewManage() *Manage {
	return &Manage{
		StudentList:     make([]Student, 0),
		CourseList:      make([]Course, 0),
		StudentScoreMap: make(map[int]Score),
		menuNum:         nil,
	}
}

func (manage Manage) showMenu() {
	fmt.Println("=======================")
	fmt.Println("=====学生信息管理系统=====")
	fmt.Println("=======================")
	fmt.Println("=====  1.学生列表  =====")
	fmt.Println("=====  2.添加学生  =====")
	fmt.Println("=====  3.编辑学生  =====")
	fmt.Println("=====  4.删除学生  =====")
	fmt.Println("=====  5.学生成绩  =====")
	fmt.Println("=======================")
}

func (manage *Manage) selectMenu() {
	orderNum := 0
	fmt.Print("请输入菜单序号: ")
	fmt.Scanln(&orderNum)
	manage.menuNum = &orderNum
}

func (manage *Manage) running() {
	menuNum := *manage.menuNum
	if menuNum == 1 {
		studentList(manage)
	} else if menuNum == 2 {
		addStudent(manage)
	} else if menuNum == 3 {
		fmt.Println("您选择的菜单又问题！")
	} else if menuNum == 4 {
		fmt.Println("您选择的菜单又问题！")
	} else if menuNum == 5 {
		fmt.Println("您选择的菜单又问题！")
	} else {
		fmt.Println("您选择的菜单又问题！")
	}
	manage.Run()
}

func addStudent(manage *Manage) {
	student := Student{}
	fmt.Print("请输入学生序号: ")
	id := 0
	fmt.Scanln(&id)
	student.Id = int8(id)
	fmt.Print("请输入学生姓名: ")
	name := ""
	fmt.Scanln(&name)
	student.Name = name
	fmt.Print("请输入学生年龄: ")
	age := 0
	fmt.Scanln(&age)
	student.Age = int8(age)
	fmt.Print("请输入学生班级: ")
	className := ""
	fmt.Scanln(&className)
	student.ClassName = className
	manage.StudentList = append(manage.StudentList, student)
	fmt.Println("添加成功！")
}

func studentList(manage *Manage) {
	if len(manage.StudentList) == 0 {
		fmt.Println("目前没有学生信息！")
	} else {
		fmt.Println("序号\tID\t姓名\t年龄\t班级\t")
		for i, student := range manage.StudentList {
			fmt.Printf("%d\t%d\t%s\t%d\t%s\t\n", i+1, student.Id, student.Name, student.Age, student.ClassName)
		}
	}
}

func (manage Manage) Run() {
	manage.showMenu()
	manage.selectMenu()
	manage.running()
	fmt.Println("******************************")
}
