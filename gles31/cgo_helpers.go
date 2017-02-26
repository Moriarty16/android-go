// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sun, 26 Feb 2017 06:20:47 MSK.
// By https://git.io/c-for-go. DO NOT EDIT.

package gles31

/*
#cgo LDFLAGS: -lGLESv3
#include <GLES3/gl31.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// unpackPCharString represents the data from Go string as *C.GLchar and avoids copying.
func unpackPCharString(str string) (*C.GLchar, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.GLchar)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

type stringHeader struct {
	Data uintptr
	Len  int
}

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// packPUbyteString creates a Go string backed by *C.GLubyte and avoids copying.
func packPUbyteString(p *C.GLubyte) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = uintptr(unsafe.Pointer(p))
		for *p != 0 {
			p = (*C.GLubyte)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
	}
	return
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(unsafe.Pointer(h.Data)), C.int(h.Len))
}

// allocPCharMemory allocates memory for type *C.GLchar in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPCharValue = unsafe.Sizeof([1]*C.GLchar{})

const sizeOfPtr = unsafe.Sizeof(&struct{}{})

// unpackArgSString transforms a sliced Go data structure into plain C format.
func unpackArgSString(x []string) (unpacked **C.GLchar, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.GLchar) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPCharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.GLchar)(unsafe.Pointer(h0))
	for i0 := range x {
		v0[i0], _ = unpackPCharString(x[i0])
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.GLchar)(unsafe.Pointer(h.Data))
	return
}

// packSString reads sliced Go data structure out from plain C format.
func packSString(v []string, ptr0 **C.GLchar) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.GLchar)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = packPCharString(ptr1)
	}
}

// packPCharString creates a Go string backed by *C.GLchar and avoids copying.
func packPCharString(p *C.GLchar) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = uintptr(unsafe.Pointer(p))
		for *p != 0 {
			p = (*C.GLchar)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
	}
	return
}
