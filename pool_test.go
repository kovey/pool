package pool

import (
	"context"
	"fmt"
	"testing"

	"github.com/kovey/pool/object"
)

type objTest struct {
	*object.Object
}

func (o *objTest) Show() {
	fmt.Println("this is obj test")
}

func TestPools(t *testing.T) {
	pool := NewPool("test", "obj", func() any {
		return &objTest{Object: object.NewObject("test", "obj")}
	})

	Reg(pool)
	obj := Get[*objTest]("test", "obj", context.Background())
	obj.Show()
	Put(obj)
}
