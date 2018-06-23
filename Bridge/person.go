/*
  Bridge 桥接模式：
        将抽象部分与它的实现部分分离，使它们都可以独立地变化
 个人想法：组合/聚合复用原则
 作者：   HCLAC
 日期：   20170306
*/

package bridge

import (
	"fmt"
)

type Person interface {
	SetCloth(c Clothes)error
	Dress()error
	Undress()error
}

type Worker struct{
	cloth Clothes
	name  string
}

func NewWorker() Person{
	p := Worker{}
	p.name = "worker"
	return &p
}

func (self *Worker)SetCloth(c Clothes)error {
	self.cloth = c
	return nil
}

func (self *Worker)Dress()error {
	fmt.Println(self.name + "--dress--" + self.cloth.GetName())
	return nil
}

func (self *Worker)Undress()error {
	fmt.Println(self.name + "--undress--" + self.cloth.GetName())
	return nil
}

type Student struct{
	cloth Clothes
	name  string
}

func NewStudent() Person{
	p := Worker{}
	p.name = "student"
	return &p
}

func (self *Student)SetCloth(c Clothes)error {
	self.cloth = c
	return nil
}

func (self *Student)Dress()error {
	fmt.Println(self.name + "--dress--" + self.cloth.GetName())
	return nil
}

func (self *Student)Undress()error {
	fmt.Println(self.name + "--undress--" + self.cloth.GetName())
	return nil
}