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

func (o *objTest) Reset() {
	o.Object.Reset()
	fmt.Println("put to pool")
}

func TestPools(t *testing.T) {
	pool := NewPool("test", "obj", func() any {
		return &objTest{Object: object.NewObject("test", "obj")}
	})

	Reg(pool)
	ctx := NewContext(context.Background())
	obj := ctx.Get("test", "obj").(*objTest)
	obj.Show()
	ctx.Drop()
}
