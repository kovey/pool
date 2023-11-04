package object

import (
	"context"
	"fmt"
)

type ObjectInterface interface {
	Reset()
	Init(ctx context.Context)
	FullName() string
}

type Object struct {
	Ctx       context.Context
	_fullName string
}

func NewObject(namespace, name string) *Object {
	return &Object{_fullName: fmt.Sprintf("%s.%s", namespace, name)}
}

func (o *Object) Reset() {
	o.Ctx = nil
}

func (o *Object) Init(ctx context.Context) {
	o.Ctx = ctx
}

func (o *Object) FullName() string {
	return o._fullName
}
