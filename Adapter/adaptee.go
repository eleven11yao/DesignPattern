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

type Pig struct {
	name string
}

func NewPig() Pig {
	p := Pig{"Pig"}
	return p
}

func (self *Pig) Run() string {
	return self.name + "---run-on-the-groud"
}

/*------------------------------------------------------------*/

type Eagle struct {
	name string
}

func NewEagle() Eagle {
	p := Eagle{"Eagle"}
	return p
}

func (self *Eagle) Fly() string {
	return self.name + "---fly-in-the-sky"
}

