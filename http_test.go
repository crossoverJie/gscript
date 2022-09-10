package gscript

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	s := struct {
		Name string `json:"name"`
	}{Name: "abc"}
	marshal, _ := json.Marshal(s)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Fprintf(w, string(marshal))
}

type h func(http.ResponseWriter, *http.Request)

func createHandle() []h {
	var hs []h
	//hs = append(hs, )
	return hs
}

func TestHttp(t *testing.T) {
	http.HandleFunc("/", handler)
	//err := http.ListenAndServe(":8000", nil)
	//if err != nil {
	//	fmt.Printf("http server failed, err:%v\n", err)
	//	return
	//}
}

func TestHttp1(t *testing.T) {
	script := `
class Person{
    string name;
}
func (HttpContext) handle (HttpContext ctx){
    Person p = Person();
    p.name = "abc";
    println("p.name=" + p.name);
    println("ctx=" + ctx);
    ctx.JSON(200, p);
}

func (HttpContext) handle1 (HttpContext ctx){
    Person p = Person();
    p.name = "def";
    println("p.name=" + p.name);
    println("ctx=" + ctx);
    ctx.JSON(200, p);
}
func (HttpContext) handle2 (HttpContext ctx){
	string local = getCurrentTime("Asia/Shanghai","2006-01-02 15:04:05");
	println(local);
    string html =^
<html>
    <title>hello</title>
    <h1>current ^+ local +^</h1>
    <p>hahaha</p>
</html>
^;
    ctx.HTML(200, html);
}
httpHandle("get", "/p", handle);
httpHandle("get", "/p/1", handle1);
httpHandle("get", "/p/2", handle2);
httpRun(":8000");
`
	NewCompiler().Compiler(script)
}

func TestDate(t *testing.T) {
	s := time.Now().String()
	fmt.Println(s)
}
