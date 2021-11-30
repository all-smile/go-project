package counters

// alertCounter 是一个未公开的类型，用于保存告警次数
type alertCounter int

// 工厂函数
// New 创建并返回一个未公开的 alertCounter 类型的值
func New(value int) alertCounter {
	return alertCounter(value)
}
