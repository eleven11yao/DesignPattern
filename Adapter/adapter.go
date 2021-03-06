/*
 Adapter 适配器模式：
        将一个类的接口转换成客户端希望的另一个接口。
		适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作
 个人想法：代理和适配器：代理和代理的对象接口一致，客户端不知道代理对象，
        而适配器是客户端想要适配器的接口，适配器对象的接口和客户端想要的不一样，
		适配器将适配器对象的接口封装一下，改成客户端想要的接口
 作者：   HCLAC
 日期：   20170306
*/

package adapter

type Adapter interface {
	Start()string
}

type PigAdapter struct {
	pig Pig
}

func NewPigAdapter() Adapter {
	p := PigAdapter{}
	p.pig = NewPig()
	return &p
}

func (self *PigAdapter) Start()string {
	return self.pig.Run()
}

type EagleAdapter struct {
	eagle Eagle
}

func NewEagleAdapter() Adapter {
	p := EagleAdapter{}
	p.eagle = NewEagle()
	return &p
}

func (self *EagleAdapter) Start()string {
	return self.eagle.Fly()
}