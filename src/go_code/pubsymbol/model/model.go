package model

// 公开 Student 结构体变量
type student struct {
	Name  string
	Age   int
	score float64 // 未公开的字段，怎么在别包获取？
}

func NewStudent(name string, age int) *student {
	return &student{
		Name: name,
		Age:  age,
	}
}

// 创建一个对外公开的方法，返回其私有的字段 score
func (stu *student) GetScore() float64 {
	return stu.score
}
