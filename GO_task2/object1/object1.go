/*
	定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。

在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
- 考察点 ：接口的定义与实现、面向对象编程风格。
*/
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
} //定义Shape接口，创建两种方法。不管是面积还是周长，都可以被归到形状的范畴

type Rectangle struct {
	width  float64
	height float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func PrintShape(s Shape) {
	TypeName := fmt.Sprintf("%T", s)[5:] //打印s内变量的具体类型，把 main.Rectangle 截成 Rectangle
	fmt.Printf("%v的面积是%.4f，周长是%.4f\n", TypeName, s.Area(), s.Perimeter())
} //调用接口

func main() {

	rect := Rectangle{width: 3, height: 4} //创建矩形实例
	circ := Circle{radius: 2.5}            //创建圆形实例
	shapes := []Shape{rect, circ}          // 创建一个接口切片，存放不同类型但都实现了 Shape 接口的对象
	for _, s := range shapes {
		PrintShape(s)
	}
}
