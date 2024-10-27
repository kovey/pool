package object

import (
	"context"
	"fmt"
)

type CtxInterface interface {
	context.Context
	Get(namespace, name string) ObjectInterface
	GetNoCtx(namespace, name string) ObjNoCtxInterface
	Parent() context.Context
}

type ObjectInterface interface {
	Reset()
	Init(ctx CtxInterface)
	FullName() string
}

type Object struct {
	Ctx       CtxInterface
	Context   context.Context
	_fullName string
}

func NewObject(namespace, name string) *Object {
	return &Object{_fullName: fmt.Sprintf("%s.%s", namespace, name)}
}

func (o *Object) Reset() {
	o.Ctx = nil
	o.Context = nil
}

func (o *Object) Init(ctx CtxInterface) {
	o.Ctx = ctx
	o.Context = ctx.Parent()
}

func (o *Object) FullName() string {
	return o._fullName
}
