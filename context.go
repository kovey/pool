package pool

import (
	"context"

	"github.com/kovey/pool/object"
)

const (
	ctx_namespace = "ko.pool.context"
	ctx_name      = "Context"
)

func init() {
	poolsNoCtx.Reg(NewPoolNoCtx(ctx_namespace, ctx_name, func() any {
		return &Context{ObjNoCtx: object.NewObjNoCtx(ctx_namespace, ctx_name)}
	}))
}

func NewContext(parent context.Context) *Context {
	ctx := GetNoCtx[*Context](ctx_namespace, ctx_name)
	ctx.Context = parent
	return ctx
}

type Context struct {
	*object.ObjNoCtx
	context.Context
	noCtxObjs []object.ObjNoCtxInterface
	objs      []object.ObjectInterface
}

func (c *Context) Parent() context.Context {
	return c.Context
}

func (c *Context) Get(namespace, name string) object.ObjectInterface {
	obj := pools.Get(namespace, name, c)
	c.objs = append(c.objs, obj)
	return obj
}

func (c *Context) GetNoCtx(namespace, name string) object.ObjNoCtxInterface {
	obj := poolsNoCtx.Get(namespace, name)
	c.noCtxObjs = append(c.noCtxObjs, obj)
	return obj
}

func (c *Context) Reset() {
	for _, obj := range c.objs {
		pools.Put(obj)
	}
	for _, obj := range c.noCtxObjs {
		poolsNoCtx.Put(obj)
	}

	c.objs = nil
	c.noCtxObjs = nil
	c.Context = nil
}

func (c *Context) Drop() {
	poolsNoCtx.Put(c)
}
