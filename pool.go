package pool

import (
	"fmt"
	"sync"

	"github.com/kovey/pool/object"
)

type PoolInterface interface {
	Put(object.ObjectInterface)
	Get(object.CtxInterface) object.ObjectInterface
	Name() string
}

type Pool struct {
	p    *sync.Pool
	name string
}

func NewPool(namespace, name string, new func() any) *Pool {
	return &Pool{name: fmt.Sprintf("%s.%s", namespace, name), p: &sync.Pool{New: new}}
}

func (p *Pool) Name() string {
	return p.name
}

func (p *Pool) Put(obj object.ObjectInterface) {
	obj.Reset()
	p.p.Put(obj)
}

func (p *Pool) Get(ctx object.CtxInterface) object.ObjectInterface {
	val, ok := p.p.Get().(object.ObjectInterface)
	if !ok {
		return nil
	}

	val.Init(ctx)
	return val
}

type Pools struct {
	pools map[string]PoolInterface
}

func NewPools() *Pools {
	return &Pools{pools: make(map[string]PoolInterface)}
}

func (p *Pools) Reg(pool PoolInterface) {
	p.pools[pool.Name()] = pool
}

func (p *Pools) Put(obj object.ObjectInterface) {
	if pool, ok := p.pools[obj.FullName()]; ok {
		pool.Put(obj)
	}
}

func (p *Pools) Get(namespace, name string, ctx object.CtxInterface) object.ObjectInterface {
	if pool, ok := p.pools[fmt.Sprintf("%s.%s", namespace, name)]; ok {
		return pool.Get(ctx)
	}

	return nil
}

var pools = NewPools()

func Reg(pool PoolInterface) {
	pools.Reg(pool)
}

func Default(namespace, name string, new func() any) {
	Reg(NewPool(namespace, name, new))
}

func Put[T object.ObjectInterface](obj T) {
	pools.Put(obj)
}

func Get[T object.ObjectInterface](namespace, name string, ctx object.CtxInterface) T {
	return pools.Get(namespace, name, ctx).(T)
}

func GetObject(namespace, name string, ctx object.CtxInterface) object.ObjectInterface {
	return pools.Get(namespace, name, ctx)
}
