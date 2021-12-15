package monster

import (
	"testing"
)

// 测试用例函数
func TestStore(t *testing.T) {
	monster := Monster{
		Name:  "牛牛",
		Age:   12,
		Skill: "顶！",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()错误， 实际为：%v, 期望为：%v", res, true)
	} else {
		t.Logf("monster.Store()成功， 实际为：%v, 期望为：%v", res, true)
	}
}

func TestReStore(t *testing.T) {
	var monster Monster
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.Store()错误， 实际为：%v, 期望为：%v", res, true)
	}

	// 进一步判断
	if monster.Name != "牛牛" {
		t.Fatalf("monster.Name错误， 实际为：%v, 期望为：%v", monster.Name, "牛牛")
	}

	t.Logf("monster.Store()成功， 实际为：%v, 期望为：%v", res, true)
}
