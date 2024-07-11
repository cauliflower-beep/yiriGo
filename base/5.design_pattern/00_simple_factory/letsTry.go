package main

import "fmt"

/*
	假设我们正在开发一个图形绘制程序，需要支持不同类型的图形，例如圆形和矩形。
	我们将使用简单工厂模式来创建这些图形对象，而无需直接暴露它们的创建细节给客户端。
*/

/**************simple_factory*************************/

// Shape represents a geometric shape.
type Shape interface {
	Draw()
}

// Circle represents a _circle shape.
type Circle struct{}

// Draw draws the _circle.
func (c *Circle) Draw() {
	fmt.Println("Drawing a _circle.")
}

// Rectangle represents a rectangle shape.
type Rectangle struct{}

// Draw draws the rectangle.
func (r *Rectangle) Draw() {
	fmt.Println("Drawing a rectangle.")
}

// ShapeType represents the type of shape.
type ShapeType int

const (
	CircleShape ShapeType = iota
	RectangleShape
)

// ShapeFactory is a simple factory to create shapes.
type ShapeFactory struct{}

// CreateShape creates a shape based on the given type.
func (sf *ShapeFactory) CreateShape(shapeType ShapeType) Shape {
	switch shapeType {
	case CircleShape:
		return &Circle{}
	case RectangleShape:
		return &Rectangle{}
	default:
		return nil
	}
}

/******************client*******************************/
func main() {
	factory := &ShapeFactory{}

	// Create a _circle
	circle := factory.CreateShape(CircleShape)
	circle.Draw()

	// Create a rectangle
	rectangle := factory.CreateShape(RectangleShape)
	rectangle.Draw()
}
