package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var template string = `{"jsonrpc": "1.0", "id":"", "method": "%s", "params": []}` 

type RPCResponse struct {
	Result int `json:"result"` 
	Error string `json:"error"` 
	Id string `json:"id"`
}

func main() {
	http.HandleFunc("/blockCount", blockCountHandler)
	http.ListenAndServe(":5000", nil)
}

func blockCountHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("POST", "http://127.0.0.1:8332", strings.NewReader(fmt.Sprintf(template, "getblockcount")))
	if err != nil {
		fmt.Println(err.Error())
	}
	req.SetBasicAuth(os.Getenv("RPCUSER"), os.Getenv("RPCPASSWD"))
	req.Header.Set("content-type", "text/plain;")	

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var data RPCResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(data.Result)

	w.Write([]byte(strconv.Itoa(data.Result)))
}