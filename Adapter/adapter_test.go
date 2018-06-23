package adapter

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	var ada Adapter
	fmt.Println("-----Start------")
	ada = NewPigAdapter()
	fmt.Println(ada.Start())
	ada = NewEagleAdapter()
	fmt.Println(ada.Start())
	fmt.Println("-----End------")
}
