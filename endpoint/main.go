//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world hello --out gen ./wit
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jamesstocktonj1/inter-component/endpoint/gen/custom/inter-component/person"
	"go.wasmcloud.dev/component/log/wasilog"
	"go.wasmcloud.dev/component/net/wasihttp"
)

var logger = wasilog.ContextLogger("record")

func init() {
	wasihttp.HandleFunc(handleRequest)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{}

	logger.Info("Fetching person...")

	res := person.GetPerson("1234")
	if res.IsErr() {
		msg["message"] = "Error fetching person"
		msg["code"] = res.Err().String()

		logger.Error("error fetching person", "error", res.Err().String())
	} else {
		p := res.OK()
		msg["name"] = p.Name
		msg["age"] = fmt.Sprintf("%d", p.Age)

		logger.Info("successfully fetched person", "person", p)
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(msg)
	w.WriteHeader(http.StatusOK)
}

func main() {}
