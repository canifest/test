package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"os"
	"strconv"

)

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

func isJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil

}

func TestList(t *testing.T) {
	var test bool

	buffer := new(bytes.Buffer)

	resp, err := http.Get("http://127.0.0.1:9993/list")
	if err != nil {
	}
	defer resp.Body.Close()
	buffer.ReadFrom(resp.Body)
	str := buffer.String()
	test = isJSONString(str)

	if test == false {
		t.Error("/list did not return a JSON formatted value")
	}

}

func TestQuit(t *testing.T) {

	buffer := new(bytes.Buffer)

	resp, err := http.Get("http://127.0.0.1:9993/pid")
	if err != nil {
	}

	defer resp.Body.Close()
	buffer.ReadFrom(resp.Body)
	pidStr := buffer.String()

	pidInt, pidErr := strconv.Atoi(pidStr)
	if pidErr != nil {
	}

	proc, err := os.FindProcess(pidInt)

	if proc != nil {
		http.Get("http://127.0.0.1:9993/quit")

	}

	resp2, err := http.Get("http://127.0.0.1:9993/pid")
	if err != nil {
	}

	if resp2 != nil {
		t.Error("Response received from Core, Quit failed.")
	}


}
