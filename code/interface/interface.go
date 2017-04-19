package main

import "fmt"


//定义一个基本的类
type Human struct{
  name  string
  age   int
  phone string
}

type Student struct{
  Human //继承 human 中的字段
  school string
  loan   float32
}

type Employee struct{
  Human
  company string
  money   float32
}


//Human 实现 Sayhi 方法
func (h Human) Sayhi(){
  fmt.Printf("hi my name is %s and my phone is %s \n", h.name,h.phone)
}

//Human 实现 Sing 方法
func (h Human) Sing(lyrics string)  {
  fmt.Println("galigaygay galigaygay ...",lyrics)
}


//student 复写 Sayhi 方法
func (sdt Student) Sayhi()  {
  fmt.Printf("my name is %s and my school is %s \n", sdt.name,sdt.school)
}

//定义一个通用的接口
type Men interface{
  Sayhi()
  Sing(lyrics string)
}



func main () {
  //实例化两种类
  mike := Student{Human{"mike",25,"18333636999"},"MIT",3.14}
  jack := Employee{Human{"jack",30,"18333636998"},"hotcast",20000}

  //定义一个接口变量
  var i Men

  //接口实现了 Student 类的方法
  fmt.Println("下面有请 mike 开始他的表演:")
  i = mike
  i.Sayhi();
  i.Sing("dongci daci")


  //接口实现了 Employee 类的方法
  fmt.Println("下面有请 mike 开始他的表演:")
  i = jack
  i.Sayhi();
  i.Sing("haha haha haha")




}
