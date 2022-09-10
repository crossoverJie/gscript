package gscript

import (
	"fmt"
	"net/http"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

type h func(http.ResponseWriter, *http.Request)

func createHandle() []h {
	var hs []h
	//hs = append(hs, )
	return hs
}

func TestHttp(t *testing.T) {
	http.HandleFunc("/", handler)
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	fmt.Printf("http server failed, err:%v\n", err)
	//	return
	//}
}

func TestHttp1(t *testing.T) {
	script := `
func void(HttpWrite, HttpRequest) handle1 (){
	
}
`
	NewCompiler().Compiler(script)
}
