package main

import (
	"testing"
)

/*
	运行 go test 进行测试，可跟上 -v 参数显示运行时间

	如果只想运行其中一个用例，例如 TestAdd 可以使用 -run 参数指定
	该参数支持通配符 *, 和部分正则表达式 例如 ^\$
	go test -run TestAdd -v
*/
func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}
}

/*************************************子测试****************************************************/
// TestMul 子测试
func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) { // 运行一个名为 "pos" 的子测试
		if Mul(2, 3) != 6 {
			t.Fatal("fail")
		}
	})
	/*
		t.Error/t.Errorf 与 t.Fatal/t.Fatalf 的区别是，前者遇错不停，还会继续执行其他的测试用例
		后者遇错即停
	*/
	t.Run("neg", func(t *testing.T) {
		if Mul(2, -3) != -6 {
			t.Fatal("fail")
		}
	})
}

/*
	TestDiv 多个子测试的场景 推荐使用如下的写法(table-driven tests)
	所有用例的数据组织在切片 cases 中，看起来就像一张表，借助循环创建子测试。
	这样写的好处有:
		1.新增用例非常简单，只需要给 cases 新增一条测试数据即可；
		2.测试代码可读性好，直观的能够看到每个子测试的参数和预期的返回值；
		3.用例失败时，报错信息的格式比较统一，测试报告易于阅读
	如果数据量较大，或是一些二进制数据，推荐使用相对路径从文件中读取
*/
func TestDiv(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 0},
		{"neg", 2, -3, 0},
		{"zero", 2, 1, 2},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Div(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got", c.A, c.B, c.Expected, ans)
			}
		})
	}
}

/*******************************************helpers*********************************************************/
type calcCase struct {
	A, B, Expected int
}

func createSubTestCase(t *testing.T, c *calcCase) {
	t.Helper() // 可以具体显示错误在哪一行
	if ans := Sub(c.A, c.B); ans != c.Expected {
		t.Fatalf("%d * %d expected %d, but %d got", c.A, c.B, c.Expected, ans)
	}
}

func TestSub(t *testing.T) {
	createSubTestCase(t, &calcCase{4, 2, 2})
	createSubTestCase(t, &calcCase{3, 4, -1})
	createSubTestCase(t, &calcCase{5, 5, 1})

}
