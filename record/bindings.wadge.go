//go:build !wasm && !wasi && !wasip1 && !wasip2 && !wasm_unknown && !tinygo.wasm

// Code generated by wadge-bindgen-go DO NOT EDIT

package main

import (
	wadge "go.wasmcloud.dev/wadge"
	"runtime"
	"unsafe"
)

const _ string = runtime.Compiler

var _ unsafe.Pointer

//go:linkname wasmimport_Log go.wasmcloud.dev/component/gen/wasi/logging/logging.wasmimport_Log
func wasmimport_Log(level0 uint32, context0 *uint8, context1 uint32, message0 *uint8, message1 uint32) {
	var __p runtime.Pinner
	defer __p.Unpin()
	if __err := wadge.WithCurrentInstance(func(__instance *wadge.Instance) error {
		return __instance.Call("wasi:logging/logging@0.1.0-draft", "log", func() unsafe.Pointer {
			ptr := unsafe.Pointer(&level0)
			__p.Pin(ptr)
			return ptr
		}(), func() unsafe.Pointer {
			ptr := unsafe.Pointer(context0)
			__p.Pin(ptr)
			return ptr
		}(), func() unsafe.Pointer {
			ptr := unsafe.Pointer(&context1)
			__p.Pin(ptr)
			return ptr
		}(), func() unsafe.Pointer {
			ptr := unsafe.Pointer(message0)
			__p.Pin(ptr)
			return ptr
		}(), func() unsafe.Pointer {
			ptr := unsafe.Pointer(&message1)
			__p.Pin(ptr)
			return ptr
		}())
	}); __err != nil {
		wadge.CurrentErrorHandler()(__err)
	}
	return
}