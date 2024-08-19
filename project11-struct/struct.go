package main

import "fmt"

type student struct {
	name string
	age  int
}

type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDreams(dreams []string) {
	// 会改变
	p.dreams = dreams
	// 下面方法不会改变
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	fmt.Println("============================")

	p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 真的想要修改 p1.dreams
	data[1] = "不睡觉"
	fmt.Println(p1.dreams)

	fmt.Println("============================")

	serialize()

	fmt.Println("============================")

	sm := &StudentManager{}

	// 添加学生
	sm.AddStudent(StudentInfo{ID: 1, Name: "张三", Age: 20, Grade: 85.5, Subjects: []string{"数学", "物理"}})
	sm.AddStudent(StudentInfo{ID: 2, Name: "李四", Age: 21, Grade: 90.0, Subjects: []string{"化学", "生物"}})

	// 展示学生列表
	sm.StudentsList()

	// 编辑学生信息
	sm.EditStudent(1, StudentInfo{ID: 1, Name: "张三", Age: 26, Grade: 88.0, Subjects: []string{"数学", "物理", "英语"}})

	// 展示学生列表
	sm.StudentsList()

	// 删除学生
	sm.DeleteStudent(2)

	// 展示学生列表
	sm.StudentsList()
}
