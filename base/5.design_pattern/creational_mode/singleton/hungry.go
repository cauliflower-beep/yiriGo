package main


type hungry struct {}	// 这里single定义为私有，防止包外new创建

var instance_hungry *hungry

/*
初始化顺序：
包级别变量初始化->init()->main()
 */

// 构造函数为私有函数，外部无法访问，自然也就无法生成实例
func init(){
	instance_hungry = &hungry{}
}

func GetHungry()*hungry{
	return instance_hungry
}

// 饿汉模式有个问题：实例instance没使用也会自动创建，造成资源消耗。此时懒汉模式登场