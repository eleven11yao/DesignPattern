package flyweight

import (
	"testing"
)

func TestFlyweight(t *testing.T) {
	ff := NewFlyweightFactory()

	fya := ff.GetFlyweight("a")
	fya.Show()
	fya.Operation(1)
	fya.Show()

	fyb := ff.GetFlyweight("b")
	fyb.Show()
	fyb.Operation(2)
	fyb.Show()

	fyc := ff.GetFlyweight("c")
	fyc.Show()
	fyc.Operation(3)
	fyc.Show()

	fyd := ff.GetFlyweight("d")
	fyd.Show()
	fyd.Operation(4)
	fyd.Show()

	fyu := ff.GetFlyweight("u")
	fyu.Show()
	fyu.Operation(5)
	fyu.Show()

	fyx := ff.GetFlyweight("b")
	fyx.Show()
	fyx.Operation(8)
	fyx.Show()
}
