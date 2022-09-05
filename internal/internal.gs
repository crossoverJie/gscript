// return string[] length
int len(string[] a){}

// return int[] length
int len(int[] a){}

// return float[] length
int len(float[] a){}

// return bool[] length
int len(bool[] a){}

int hash(string s){}
int hash(int s){}
int hash(float s){}
int hash(bool s){}

println(){}
assertEqual(){}


class EntryString{
    string key,value;
    EntryString next;
    EntryString(string k, string v, EntryString n){
        key=k;
        value=v;
        next=n;
    }
}
class MapString{
    EntryString[] table = [16]{};
    int size=0;

    put(string key, string value){
        // todo crossoverJie 扩容
        if(key==""){
            return;
        }
        int hashcode = hash(key);
        int i = hashcode % len(table) ;
        EntryString e = table[i];
        bool write = true;
        if (e != nil){
            EntryString tempEntry = e;
            for (tempEntry != nil) {
                int currentHash = hash(tempEntry.key);
                if (currentHash == hashcode && key == tempEntry.key){
                    tempEntry.value= value;
                    write =false;
                    return;
                }
                tempEntry = tempEntry.next;
            }
            table[i] = EntryString(key, value, e);
            if (write){
                size++;
            }

        } else {
            // 参考 jdk1.7 链表头插法
            table[i] = EntryString(key,value,nil);
            size++;
        }

    }

    string get(string key){
        if(key==""){
            return;
        }
        int hashcode = hash(key);
        int i = hashcode % len(table) ;
        EntryString e = table[i];
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
    String value;
    LinkedNode next;
    LinkedNode(string v){
        value = v;
        next = n;
    }
}
class LinkedList{
    LinkedNode first, last;
    int size=0;

    add(string v){
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