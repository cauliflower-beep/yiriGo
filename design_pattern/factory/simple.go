package main

import "fmt"

/*
	简单工厂模式：
	定义一个工厂类，它可以根据参数的不同返回不同类的实例，被创建的实例通常都具有共同的父类。
	因为在简单工厂模式中用于创建实例的方法是静态（static）方法，
	因此简单工厂模式又被称为静态工厂方法（Static Factory Method）模式，它属于类创建型模式。

	简单工厂需要：
	工厂结构体、产品接口、产品结构体

	优点：
	工厂类是整个工厂模式的核心，我们只需要传入给定的信息，就可以创建所需实例。
	在多人协作的时候，无需知道对象之间的内部依赖，可以直接创建，有利于整个软件体系结构的优化。

	缺点:
	工厂类中包含了所有实例的创建逻辑，一旦这个工厂类出现问题，所有实例都会受到影响.
	并且，工厂类中生产的产品都基于一个共同的接口，一旦要添加不同种类的产品，这就会增加工厂类的复杂度.
	将不同种类的产品混合在一起，违背了单一职责，系统的灵活性和可维护性都会降低;
	并且当新增产品的时候，必须要修改工厂类，违背了『系统对扩展开放，对修改关闭』的原则。
*/

// 创建一个饺子店工厂结构体，和饺子类的接口。该工厂的其中一个方法用来生产不同口味的饺子，如韭菜的猪肉馅的。
//type DumplingsShop struct {
//	//Generate(t string) *Dumplings
//}
//type Dumplingsinterface interface {
//	create()
//}
//
////创建肉馅和韭菜馅的饺子结构体，并且实现对应接口的方法。
//type DumplingsMeat struct{}
//
//func (*DumplingsMeat) create() {
//	fmt.Println("DumplingsMeat create")
//}
//
//type DumplingsChives struct{}
//
//func (*DumplingsChives) create() {
//	fmt.Println("DumplingsChives create")
//}
//
//func (*DumplingsShop) Create(stuff string) *Dumplings {
//	switch stuff {
//	case "meat":
//		return new(DumplingsMeat)
//	case "chives":
//		return new(DumplingsChives)
//	default:
//		return nil
//	}
//}

// 创建一个文章接口，实现阅读和写作功能
type Article interface {
	ReadArticle() string
	WriteArticle(content string) string
}

// 创建中文书和英文书 两个"类"，并分别实现文章接口的两个方法。
type ChineseArticle struct{}
type EnglishArticle struct{}

func (c *ChineseArticle) ReadArticle() string {
	return "这是中文书"
}
func (c *EnglishArticle) ReadArticle() string {
	return "this's English book"
}
func (c *ChineseArticle) WriteArticle(contents string) string {
	return "天气真好呀"
}
func (c *EnglishArticle) WriteArticle(contents string) string {
	return "how sunny today!"
}

// 使用断言方式分配不同的参数，应该需要使用什么样的方法去实例化具体的类
func SFactory(lan string) (art Article) {
	switch lan {
	case "Chinese":
		art = &ChineseArticle{}
	case "English":
		art = &EnglishArticle{}
	default:
		art = &ChineseArticle{}
	}
	return art
}

// 使用
func main() {
	//var stuff string
	//dumplingFactory := DumplingsShop{}
	//stuff = "meat"
	//meat := dumplingFactory.Create(stuff)//返回肉馅饺子对象
	//meat.create()
	//stuff = "chives"
	//chives := dumplingFactory.Create(stuff) //返回韭菜馅饺子对象
	//chives.create()

	a := SFactory("Chinese") //这样，我实例化中文书还是英文书都可以了
	fmt.Println(a.ReadArticle())
}
