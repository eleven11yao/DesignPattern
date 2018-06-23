/*
  Bridge 桥接模式：
        将抽象部分与它的实现部分分离，使它们都可以独立地变化
 个人想法：组合/聚合复用原则
 作者：   HCLAC
 日期：   20170306
*/

package bridge

type Clothes interface {
	GetName()string
}

type Uniform struct {
	name string
}

func NewUniform() Clothes {
	t := Uniform{}
	t.name = "Uniform"
	return &t
}

func (self *Uniform) GetName() string {	
	return self.name
}


type Shirt struct {
	name string
}

func NewShirt() Clothes {
	t := Shirt{}
	t.name = "Shirt"
	return &t
}

func (self *Shirt) GetName() string {	
	return self.name
}
