int num(int x,int y){
	if (y==1 || y ==x) {
		return 1;
	}
    int v1 = num(x - 1, y - 1);
    int v2 = num(x - 1, y);
	int c = v1+v2;
    // int c = num(x - 1, y - 1)+num(x - 1, y);
	return c;
}

printTriangle(int row){
	for (int i = 1; i <= row; i++) {
        for (int j = 1; j <= row - i; j++) {
           print(" ");
        }
        for (int j = 1; j <= i; j++) {
            print(num(i, j) + " ");
        }
        println("");
    }
}

printTriangle(7);
