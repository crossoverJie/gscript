void fib(){
    int a = 0;
    int b = 1;
    int fibonacci(){
        int c = a;
        a = b;
        b = a+c;
        return c;
    }
    return fibonacci;
}

func int() f = fib();

for (int i = 0; i < 10; i++){
    print(f()+" ");
}