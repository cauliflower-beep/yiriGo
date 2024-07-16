package main

import "fmt"

// AModuleAPI ...
type AModuleAPI interface {
	TestA() string
}

type aModuleAPI struct{}

func (*aModuleAPI) TestA() string {
	return "A module is running"
}
func NewAModuleAPI() AModuleAPI {
	return &aModuleAPI{}
}

// BModuleAPI ...
type BModuleAPI interface {
	TestB() string
}

type bModuleAPI struct{}

func (*bModuleAPI) TestB() string {
	return "B module is running"
}
func NewBModuleAPI() BModuleAPI {
	return &bModuleAPI{}
}

// API facade模块的外观接口，大部分代码使用此接口简化对facade类的访问。
type API interface {
	Test() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%S\n%s", aRet, bRet)
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}
