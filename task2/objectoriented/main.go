package main

import (
	"fmt"
	"math"
)

/*
	题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
		  在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
	考察点 ：接口的定义与实现、面向对象编程风格。

	题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
		  为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	考察点 ：组合的使用、方法接收者。
*/

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  int
	Height int
}

func (obj Rectangle) Area() float64 {
	return float64(obj.Width * obj.Height)
}

func (obj Rectangle) Perimeter() float64 {
	return float64(obj.Width*obj.Height) / 2
}

type Circle struct {
	Radius float64
}

func (obj Circle) Area() float64 {
	return math.Pi * obj.Radius * obj.Radius
}

func (obj Circle) Perimeter() float64 {
	return 2 * math.Pi * obj.Radius
}

type Person struct {
	Age  int
	Name string
}

type Employee struct {
	Person     Person
	EmployeeID int
}

func (employee Employee) PrintInfo() {
	fmt.Printf("This is %s, age is %d, id is %d \n", employee.Person.Name, employee.Person.Age, employee.EmployeeID)
}

func main() {
	shapes := []Shape{
		Rectangle{2, 4},
		Circle{2.5},
	}
	for i := 0; i < len(shapes); i++ {
		fmt.Printf("面积：%v, 周长： %v \n", shapes[i].Area(), shapes[i].Perimeter())
	}

	employee := Employee{
		Person{12, "xiaoming"},
		1008,
	}
	employee.PrintInfo()
}
