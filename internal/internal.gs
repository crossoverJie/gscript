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


class EntryString{
    string key,value;
    EntryString next;
    EntryString(string k, string v){
        key=k;
        value=v;
    }
}
class MapString{
    EntryString[] table = [16]{};
    int size=0;

    put(string key,string value){
        if(key==""){
            return;
        }
        int hashcode = hash(key);
        int i = hashcode % len(table) ;
        EntryString e = table[i];
        // for (e.next != nil){
        //     int tempHash = hash(e.key);
        //     if (hashcode == tempHash){
        //         tempEntry = EntryString(key,value);
        //         e
        //     }else {
        //
        //     }
        //
        //     e = e.next;
        // }
        e = EntryString(key,value);
        table[i]=e;
        size++;
    }

string get(string key){
    if(key==""){
        return;
    }
    int hashcode = hash(key);
    int i = hashcode % len(table) ;
    EntryString e = table[i];
    //print("get e=" + e);
    return e.value;
}

    int getSize(){
        return size;
    }
}


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
            print(start.value);
            start = start.next;
        }
    }

    int getSize(){
        return size;
    }
}