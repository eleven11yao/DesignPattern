package bridge

import (
	"fmt"
	"testing"
)

//func TestBridge(t *testing.T) {
//	sa := SoftwareA{Software{"a"}}
//	sb := SoftwareB{Software{"b"}}
//
//	pa := NewPhoneA("pa")
//	pb := NewPhoneB("pb")
//
//	pa.setSoft(&sa)
//	pa.Run()
//
//	pb.setSoft(&sb)
//	pb.Run()
//
//	fmt.Println()
//	p := TSoftware{&sb}
//	p.Run()
//}

func TestBridgeStep1(t *testing.T) {
	c := NewUniform()
	s := NewShirt()	
	w := NewWorker()
	
	fmt.Println("-----Start------")
	w.SetCloth(c)
	w.Dress()
	w.Undress()
	w.SetCloth(s)
	w.Dress()
	w.Undress()
	fmt.Println("-----End------")

	fmt.Println("-----Start------")
	stu := NewStudent()
	stu.SetCloth(c)
	stu.Dress()
	stu.Undress()
	fmt.Println("-----End------")

}