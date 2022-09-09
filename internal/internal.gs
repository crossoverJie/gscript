// return array length
int len(any[] a){}

// return hashcode
int hash(any s){}

println(){}
assertEqual(){}
append(){}

// return JSON string
string JSON(any a){}
// JSON query with path
any JSONGet(string json, string path){}

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
        bool write = true;
        if (e != nil){
            Entry tempEntry = e;
            for (tempEntry != nil) {
                int currentHash = hash(tempEntry.key);
                if (currentHash == hashcode && key == tempEntry.key){
                    tempEntry.value= value;
                    write =false;
                    return;
                }
                tempEntry = tempEntry.next;
            }
            table[i] = Entry(key, value, e);
            if (write){
                size++;
            }

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
        return e.value;
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
        next = n;
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