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
