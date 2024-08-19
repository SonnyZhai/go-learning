package main

import "fmt"

type StudentInfo struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Grade    float64  `json:"grade"`
	Subjects []string `json:"subjects"`
}

type StudentManager struct {
	Students []StudentInfo
}

func (sm *StudentManager) AddStudent(s StudentInfo) {
	sm.Students = append(sm.Students, s)
	fmt.Println("添加学生成功")
}

func (sm *StudentManager) DeleteStudent(id int) {
	for i, stu := range sm.Students {
		if stu.ID == id {
			sm.Students = append(sm.Students[:i], sm.Students[i+1:]...)
			fmt.Println("删除学生成功")
			return
		}
	}
	fmt.Println("未找到对应的学生ID，删除学生失败")
}

func (sm *StudentManager) EditStudent(id int, newStu StudentInfo) {
	for i, stu := range sm.Students {
		if stu.ID == id {
			sm.Students[i] = newStu
			fmt.Println("修改学生成功")
			return
		}
	}
	fmt.Println("未找到对应的学生ID，修改学生失败")
}

func (sm *StudentManager) StudentsList() {
	fmt.Println("学生列表:")
	for _, s := range sm.Students {
		fmt.Printf("ID: %d, 姓名: %s, 年龄: %d, 分数: %.2f, 科目: %v\n", s.ID, s.Name, s.Age, s.Grade, s.Subjects)
	}
}
