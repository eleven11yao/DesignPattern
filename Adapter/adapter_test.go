package adapter

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	fmt.Println("-----Start------")
	fmt.Println(NewPigAdapter().Start())
	fmt.Println(NewEagleAdapter().Start())
	fmt.Println("-----End------")
}
