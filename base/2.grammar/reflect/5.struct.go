package main

import (
	"fmt"
	"reflect"
)

type Robot struct {
	Id    int    `json:"id"`
	speed int    `json:"speed"`
	force int    `json:"force"`
	Motto string `json:"motto"`
}

func (r Robot) GetInfo() string {
	return fmt.Sprintf("this robot's id is %v,speed is %v,force is %v,motto is %v", r.Id, r.speed, r.force, r.Motto)
}

func (r *Robot) SetInfo(id, speed, force int, motto string) {
	r.Id = id
	r.speed = speed
	r.force = force
	r.Motto = motto
}

// StructField 利用反射获取结构体字段
func StructField(r interface{}) {
	t := reflect.TypeOf(r)  // 类型变量
	v := reflect.ValueOf(r) // 值变量
	// 两个判断条件保证了参数为值传递或者地址传递均有效
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("r is not a struct!")
		return
	}
	// 1. 通过类型变量里面的Field可以获取结构体字段
	/*
		注意方法接收者 Robot 和 *Robot 是不同的！！！
	*/
	field0 := t.Field(0)
	fmt.Printf("%#v \n", field0)
	fmt.Println(field0)
	fmt.Println(field0.Name)            // 字段名称
	fmt.Println(field0.Type)            // 字段类型
	fmt.Println(field0.Tag.Get("json")) // 字段标签，因为标签可以定义很多个，所以要通过这种形式来获取
	// 2. 通过类型变量里面的FieldByName可以获取结构体的字段
	field1, ok := t.FieldByName("speed")
	if !ok {
		fmt.Println("r has not field which name speed.")
	}
	fmt.Println(field1)
	fmt.Println(field1.Index)           // 字段索引
	fmt.Println(field1.Type)            // 字段类型
	fmt.Println(field1.Tag.Get("json")) // 字段标签
	// 3. 通过类型变量里面的NumFiled获取到该结构体有几个字段
	num := t.NumField()
	/*
		NumField中已经做了类型检查，当发现调用者不是结构体时，就会抛出panic
	*/
	fmt.Println("{type's fieldNum:%d}", num)
	// 4. 通过值变量获取结构体属性对应的值
	value0 := v.Field(0)
	fmt.Println(value0)
	fieldNum := v.NumField()
	fmt.Println("{value's fieldNum:%d}", fieldNum)
	fmt.Println("*******************end************************")
}

// StructMethod 利用反射获取结构体的方法信息
func StructMethod(r interface{}) {
	t := reflect.TypeOf(r)
	v := reflect.ValueOf(r)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("r is not a struct")
		return
	}
	// 1. 通过类型变量里面的Method可以获取结构体的方法
	method0 := t.Method(0) // 这里的索引0和结构体方法的顺序是没有关系的，和结构体方法的ASCII有关系
	//method1,ok := t.MethodByName("GetInfo") // 通过名称获取结构体方法
	//if ok{
	//
	//}
	fmt.Println(method0)
	fmt.Println(method0.Name) // 方法名称
	fmt.Println(method0.Type) // 方法类型，并不常用
	// 2. 通过类型变量获取这个结构体有多少个方法
	fmt.Println(t.NumMethod())
	// 3. 通过"值变量"执行方法 （注意需要使用值变量，并且要注意参数） v.Method(0).Call(nil) 或者 v.MethodByName("").Call(nil)
	//info := v.Method(0).Call(nil)
	//fmt.Println(info) // 返回的是一个 []Value
	//fmt.Println(v.MethodByName("GetInfo").Call(nil))
	// 4. 执行方法传入参数 (注意需要使用“值变量”，并且要注意参数，接受的参数是[]reflect.Value的切片)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(2))
	params = append(params, reflect.ValueOf(4))
	params = append(params, reflect.ValueOf(4))
	params = append(params, v.Elem().Field(3))
	v.MethodByName("SetInfo").Call(params)
	// 5. 执行方法获取方法的值

}

// ChangeStruct 利用反射修改结构体属性
func ChangeStruct(r interface{}) {
	t := reflect.TypeOf(r)
	v := reflect.ValueOf(r)

	// 结构体是值传递，若要修改它的属性，则需要传递指针
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		fmt.Println("r is not a struct ptr")
		return
	}

	// 修改结构体属性的值,注意待修改的属性需要是导出字段
	id := v.Elem().FieldByName("Id")
	id.SetInt(0)
	fmt.Println("***************end*****************")
}
func main() {
	sam := Robot{
		Id:    1,
		speed: 7,
		force: 7,
		Motto: "流水在触碰到底处时才会释放活力",
	}
	StructField(sam)
	StructMethod(&sam)
	fmt.Println(sam.GetInfo())

	ChangeStruct(&sam)
	fmt.Println(sam.GetInfo())

}
