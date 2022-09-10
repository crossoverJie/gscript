
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
    string queryPath = ctx.queryPath();
    println("queryPath = " + queryPath);

    // http://127.0.0.1:8000/p/2?id=100
    string id = ctx.formValue("id");
    println("id="+id);
    ctx.HTML(200, html);
}
httpHandle("get", "/p", handle);
httpHandle("get", "/p/1", handle1);
httpHandle("get", "/p/2", handle2);
httpRun(":8000");
