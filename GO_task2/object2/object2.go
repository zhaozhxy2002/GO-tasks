/*
 2. 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，

组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
- 考察点 ：组合的使用、方法接收者。
*/
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {

	fmt.Printf("👤 员工信息:\n")
	fmt.Printf("   姓名: %s\n", e.Name)
	fmt.Printf("   年龄: %d\n", e.Age)
	fmt.Printf("   工号: %s\n", e.EmployeeID)
	fmt.Println("  ---------")
}

func main() {
	Emp := Employee{
		Person:     Person{Name: "Bob", Age: 26},
		EmployeeID: "B123456"}
	Emp.PrintInfo()
}
