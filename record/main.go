//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world hello --out gen ./wit
package main

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"github.com/jamesstocktonj1/inter-component/record/gen/custom/inter-component/person"
	"go.wasmcloud.dev/component/log/wasilog"
)

var logger = wasilog.ContextLogger("record")

func init() {
	person.Exports.SetPerson = handleSetPerson
	person.Exports.GetPerson = handleGetPerson
}

func handleSetPerson(id person.PersonID, value person.Person) person.Error {
	return person.ErrorNone
}

func handleGetPerson(id person.PersonID) cm.Result[person.PersonShape, person.Person, person.Error] {
	return cm.OK[cm.Result[person.PersonShape, person.Person, person.Error]](person.Person{})
}

func main() {}
