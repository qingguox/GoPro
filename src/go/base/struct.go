package main

import (
	"encoding/json"
	. "fmt"
	"time"
)

type User struct { // 类似java中的类, 但是只能定义变量, 不能定义方案
	Name string `json:"name"` // 大写 public 小写private private是包内访问 main这个包
	Age  int    `json:"age"`  // 大写此时json序列化才能出来

	// 有json:这个序列化后, 名称就以name, age输出, 如果没有json标记, 则默认是Name
	Time int64 `json:"-"` // 标记可以忽略该字段, 不参与序列化
}

// 为User类型绑定 ToString方法, *User为指针引用可以修改传入参数的值
// 方法归属于类型, 不归属于具体的对象, 声明该类型的对象即可调用该类型的方法
func (x *User) ToString() string {
	return Sprintf("Name:%s, Age:%d", x.Name, x.Age)
}

func main() {
	var user1 User
	user1 = User{Name: "张三", Age: 20}
	PrintUser(user1)

	user2 := User{"李四", 21, time.Now().Unix()}
	PrintUser(user2)

	// 没有明确定义的 Name = ""
	user3 := User{Age: 201}
	PrintUser(user3)

	// 修改属性值
	PrintUserChange(&user3, "修改后的名字", 1001)

	// 方法
	var user1Desc = user1.ToString()
	Println(user1Desc)

	Println("json test")
	// json序列化问题
	if result, err := json.Marshal(&user2); err == nil {
		Println(string(result))
	}
}

// 这种修改是不生效的
func PrintUser(user User) {
	Printf("Name:%s, Age:%d\n", user.Name, user.Age)
}

// 只有指针才能修改具体的属性值,
func PrintUserChange(user *User, name string, age int) {
	Printf("changeBefore Name:%s, Age:%d\n", user.Name, user.Age)
	user.Name = name
	user.Age = age
	Printf("changeAfter Name:%s, Age:%d\n", user.Name, user.Age)
}
