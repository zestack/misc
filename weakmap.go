package misc

import (
	"reflect"
	"runtime"
	"sync"
)

type WeakMap struct {
	m map[uintptr]any
	l sync.RWMutex
}

func NewWeakMap() *WeakMap {
	return &WeakMap{
		m: make(map[uintptr]any),
	}
}

// Put key is weakly referenced.
func (w *WeakMap) Put(key, value any) {
	runtime.SetFinalizer(key, w.finalizer)
	ptr := reflect.ValueOf(key).Pointer()
	w.l.Lock()
	defer w.l.Unlock()
	w.m[ptr] = value
}

func (w *WeakMap) Get(key any) (any, bool) {
	ptr := reflect.ValueOf(key).Pointer()
	w.l.RLock()
	defer w.l.RUnlock()
	v, ok := w.m[ptr]
	return v, ok
}

func (w *WeakMap) Has(key any) bool {
	ptr := reflect.ValueOf(key).Pointer()
	w.l.RLock()
	defer w.l.RUnlock()
	_, ok := w.m[ptr]
	return ok
}

func (w *WeakMap) Delete(key any) {
	w.finalizer(key)
}

func (w *WeakMap) Len() int {
	w.l.RLock()
	defer w.l.RUnlock()
	v := len(w.m)
	return v
}

// TODO(hupeh): 是否需要引用次数，当次数归零时才被清理
func (w *WeakMap) finalizer(key any) {
	ptr := reflect.ValueOf(key).Pointer() // uintptr(unsafe.Pointer(&key))
	w.l.Lock()
	delete(w.m, ptr)
	w.l.Unlock()
}
