package object

import (
	"fmt"
)

type ObjNoCtxInterface interface {
	Reset()
	Init()
	FullName() string
}

type ObjNoCtx struct {
	_fullName string
}

func NewObjNoCtx(namespace, name string) *ObjNoCtx {
	return &ObjNoCtx{_fullName: fmt.Sprintf("%s.%s", namespace, name)}
}

func (o *ObjNoCtx) Reset() {
}

func (o *ObjNoCtx) Init() {
}

func (o *ObjNoCtx) FullName() string {
	return o._fullName
}
