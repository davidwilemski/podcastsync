package jsonrpc

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/gorilla/http"

	"github.com/gorilla/rpc/json"
)

// Request does nessecary encoding of params and makes the HTTP JSON RPC call
func Request(url, method string, args interface{}, reply interface{}) error {
	req, err := json.EncodeClientRequest(method, args)
	status, headers, r, err := http.DefaultClient.Post(url, map[string][]string{"Content-Type": []string{"application/json"}}, bytes.NewBuffer(req))
	if err != nil {
		return err
	}
	fmt.Println(status)
	fmt.Println(headers)
	if r == nil {
		return errors.New("blah")
	}
	defer r.Close()

	if status.Code < 200 || status.Code > 299 {
		return errors.New("blah status error")
	}

	/*	buf := new(bytes.Buffer)
		buf.ReadFrom(r)*/

	err = json.DecodeClientResponse(r, &reply)
	if err != nil {
		reply = nil
		return err
	}

	return nil
}
