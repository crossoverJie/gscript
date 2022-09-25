// return array length
int len(any[] a){}

// return hashcode
int hash(any s){}

println(any a){}

//formats according to a format specifier and writes to standard output.
printf(string format, any ...a){}

//formats according to a format specifier and returns the resulting string.
string sprintf(string format, any ...a){}

// formats using the default formats for its operands and writes to standard output.
string print(any ...a){}

assertEqual(any a1, any a2){}

// appends "v" to the end of a array "a"
append(any[] a, any v){}

// Date
string getCurrentTime(string tz, string layout){}

// Get os args.
string[] getOSArgs(){}

// return JSON string
string JSON(any a){}
// JSON query with path
any JSONGet(string json, string path){}


// http lib
// Response json
FprintfJSON(int code, string path, string json){}
// Resonse html
FprintfHTML(int code, string path, string html){}

// path (relative paths may omit leading slash)
string QueryPath(string path){}

string FormValue(string path, string key){}
class HttpContext{
    string path;
    JSON(int code, any v){
        string json = JSON(v);
        FprintfJSON(code, path, json);
    }
    HTML(int code, any v) {
        string html = v;
        FprintfHTML(code, path, html);
    }
    string queryPath() {
        string p = QueryPath(path);
        return p;
    }

    string formValue(string key){
        string v = FormValue(path, key);
        return v;
    }
}
// Bind route
httpHandle(string method, string path, func (HttpContext) handle){
    // println("path="+path);
    HttpContext ctx = HttpContext();
    handle(ctx);
}
// Run http server.
httpRun(string addr){}

class Entry{
    any key,value;
    Entry next;
    Entry(any k, any v, Entry n){
        key=k;
        value=v;
        next=n;
    }
}
class Map{
    Entry[] table = [16]{};
    int size=0;

    put(any key, any value){
        // todo crossoverJie 扩容
        if(key==""){
            return;
        }
        int hashcode = hash(key);
        int i = hashcode % len(table) ;
        Entry e = table[i];
        if (e != nil){
            Entry tempEntry = e;
            for (tempEntry != nil) {
                int currentHash = hash(tempEntry.key);
                if (currentHash == hashcode && key == tempEntry.key){
                    tempEntry.value= value;
                    return;
                }
                tempEntry = tempEntry.next;
            }
            table[i] = Entry(key, value, e);
            size++;

        } else {
            // 参考 jdk1.7 链表头插法
            table[i] = Entry(key,value,nil);
            size++;
        }

    }

    any get(any key){
        if(key==""){
            return;
        }
        int hashcode = hash(key);
        int i = hashcode % len(table) ;
        Entry e = table[i];
        for (e.next != nil){
            int currentHash = hash(e.key);
            if (key == e.key && currentHash == hashcode){
                return e.value;
            }
            e = e.next;
        }
        if (hash(e.key) == hashcode  && key == e.key){
            return e.value;
        }else {
            return nil;
        }
    }

    int getSize(){
        return size;
    }
}


// LinkedList
class LinkedNode{
    any value;
    LinkedNode next;
    LinkedNode(any v){
        value = v;
        next = nil;
    }
}
class LinkedList{
    LinkedNode first, last;
    int size=0;

    add(any v){
        if (v == ""){
            return;
        }
        LinkedNode l = LinkedNode(v);
        if (first == nil){
            first = l;
            last = l;
        } else{
            last.next = l;
            last = l;
        }
            size ++;
        }

    dump(){
        LinkedNode start = first;
        for (start != nil){
            println(start.value);
            start = start.next;
        }
    }

    int getSize(){
        return size;
    }
}