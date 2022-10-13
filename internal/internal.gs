// return array length
int len(any[] a){}

// return array cap
int cap(any[] a){}

// copies elements from a source array into a destination array.
int copy(byte[] dst, byte[] src){}

// return hashcode
int hash(any s){}

println(any a){}

//formats according to a format specifier and writes to standard output.
printf(string format, any ...a){}

//formats according to a format specifier and returns the resulting string.
string sprintf(string format, any ...a){}

// formats using the default formats for its operands and writes to standard output.
print(any ...a){}

assertEqual(any a1, any a2){}

// covert string to byte array.
byte[] toByteArray(string s){}
// covert byte array to string.
string toString(byte[] b){}
class StringBuilder{
    byte[] buf = [0]{};

    // append contens to buf, it returns the length of s
    int writeString(string s){
        byte[] temp = toByteArray(s);
        append(buf, temp);
        return len(temp);
    }
    
    // append b to buf, it returns the length of b.
    int WriteBytes(byte[] b){
        append(buf, b);
        return len(b);
    }

    grow(int n){
        if (n > 0) {
            // when there is not enough space left.
            if (cap(buf) - len(buf) < n) {
                byte[] newBuf = [len(buf), 2*cap(buf)+n]{};
                copy(newBuf, buf);
                buf = newBuf;
            }
        }   
    }

    string String(){
        return toString(buf);
    }
}

class Strings{
    string join(string[] elems, string sep){
        if (len(elems) == 0) {
            return "";
        }
        if (len(elems) == 1) {
            return elems[0];
        }
        
        byte[] bs = toByteArray(sep);
        int n = len(bs) * (len(elems) -1);
        for (int i=0; i < len(elems); i++) {
            string s = elems[i];
            byte[] bs = toByteArray(s);
            n = n + len(bs);
        }
        
        StringBuilder sb = StringBuilder();
        sb.grow(n);
        string first = elems[0];
        sb.writeString(first);

        string[] remain = elems[1:len(elems)];
        for(int i=0; i < len(remain); i++){
            sb.writeString(sep);
            string r = remain[i];
            sb.writeString(r);
        }
        return sb.String();

    }
}

// appends "v" to the end of a array "a"
append(any[] a, any v){}

// return JSON string
string JSON(any a){}

// JSON query with path
any JSONGet(string json, string path){}


// HTTP lib
// Response json
FprintfJSON(int code, string path, string json){}
// Resonse html
FprintfHTML(int code, string path, string html){}

// path (relative paths may omit leading slash)
string QueryPath(string path){}

// returns the first value for the named component of the query.
string FormValue(string path, string key){}

// returns the first value for the named component of the POST
string PostFormValue(string path, string key){}

// return request raw body
string RequestBody(string path){}

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

    string postFormValue(string key){
        string v = PostFormValue(path, key);
        return v;
    }    

    string requestBody(){
        string body = RequestBody(path);
        return body;
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

// System os api
string[] GetOSArgs(){}
string Command(string name, string ...arg){}
WriteFile(string fileName, string value, int perm){}
Remove(string fileName){}
string Getwd(){}


// system os api
class System{

    // Get os args.
    string[] getOSArgs(){
        return GetOSArgs();
    }
    string command(string name, string ...arg){
        return Command(name, arg);
    }
    // attention: perm is the decimal
    writeFile(string fileName, string value, int perm){
        WriteFile(fileName, value, perm);
    }
    // removes the named file
    remove(string fileName){
        Remove(fileName);
    }
    // returns a rooted path name corresponding to the current directory.
    string getwd(){
        return Getwd();
    }
}

// Date
string GetCurrentTime(string tz, string layout){}

int Unix(string tz){}

class DateTime{
    string getCurrentTime(string tz, string layout){
        return GetCurrentTime(tz, layout);
    }
    int unix(string tz){
        return Unix(tz);
    }
}

// Compiler api
string dumpAST(string code){}
string dumpSymbol(string code){}

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