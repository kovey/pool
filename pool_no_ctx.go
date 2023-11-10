package pool

import (
	"fmt"
	"sync"

	"github.com/kovey/pool/object"
)

type PoolNoCtxInterface interface {
	Put(object.ObjNoCtxInterface)
	Get() object.ObjNoCtxInterface
	Name() string
}

type PoolNoCtx struct {
	p    *sync.Pool
	name string
}

func NewPoolNoCtx(namespace, name string, new func() any) *PoolNoCtx {
	return &PoolNoCtx{name: fmt.Sprintf("%s.%s", namespace, name), p: &sync.Pool{New: new}}
}

func (p *PoolNoCtx) Name() string {
	return p.name
}

func (p *PoolNoCtx) Put(obj object.ObjNoCtxInterface) {
	obj.Reset()
	p.p.Put(obj)
}

func (p *PoolNoCtx) Get() object.ObjNoCtxInterface {
	val, ok := p.p.Get().(object.ObjNoCtxInterface)
	if !ok {
		return nil
	}

	val.Init()
	return val
}

type PoolNoCtxs struct {
	pools map[string]PoolNoCtxInterface
}

func NewPoolNoCtxs() *PoolNoCtxs {
	return &PoolNoCtxs{pools: make(map[string]PoolNoCtxInterface)}
}

func (p *PoolNoCtxs) Reg(pool PoolNoCtxInterface) {
	p.pools[pool.Name()] = pool
}

func (p *PoolNoCtxs) Put(obj object.ObjNoCtxInterface) {
	if pool, ok := p.pools[obj.FullName()]; ok {
		pool.Put(obj)
	}
}

func (p *PoolNoCtxs) Get(namespace, name string) object.ObjNoCtxInterface {
	if pool, ok := p.pools[fmt.Sprintf("%s.%s", namespace, name)]; ok {
		return pool.Get()
	}

	return nil
}

var poolsNoCtx = NewPoolNoCtxs()

func RegNoCtx(pool PoolNoCtxInterface) {
	poolsNoCtx.Reg(pool)
}

func DefaultNoCtx(namespace, name string, new func() any) {
	RegNoCtx(NewPoolNoCtx(namespace, name, new))
}

func PutNoCtx[T object.ObjNoCtxInterface](obj T) {
	poolsNoCtx.Put(obj)
}

func GetNoCtx[T object.ObjNoCtxInterface](namespace, name string) T {
	return poolsNoCtx.Get(namespace, name).(T)
}

func GetObjNoCtx(namespace, name string) object.ObjNoCtxInterface {
	return poolsNoCtx.Get(namespace, name)
}
