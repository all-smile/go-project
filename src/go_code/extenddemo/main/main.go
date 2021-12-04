package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}

func (stu *Student) ShowInfo() {
	fmt.Printf("名字：%v, 年龄： %v, 分数：%v \n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score float64) {
	stu.Score = score
}

// 新增方法
func (stu *Student) DoHomework() {
	fmt.Println("学生在做作业")
}

// 小学生考试
type Pupil struct {
	Student
}

// 保留 Pupil 特有方法
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试")
}

// 大学生考试
type Graduate struct {
	Student
}

// 保留 Graduate 特有方法
func (p *Graduate) testing() {
	fmt.Println("大学生正在考试")
}

func main() {
	// 小学生
	var pupil = &Pupil{}
	pupil.Name = "tom"
	pupil.Age = 6
	pupil.testing()
	pupil.SetScore(99.89)
	pupil.ShowInfo()
	pupil.DoHomework()

	// 大学生

	var graduate = &Graduate{}
	graduate.Name = "bill"
	graduate.Age = 19
	graduate.testing()
	graduate.SetScore(99.89)
	graduate.ShowInfo()
	pupil.DoHomework()
}
