package main
/*
    #include <stdio.h>
    #include <stdlib.h>
    void hello() {
        printf("Hello, World!\n");
    }
    int arr[10] = {1,4,5,6,7,3,2,3,5,2};

    void demo(){
    	for (int i=0 ; i<10 ; i++){
    		printf(arr[i]);
    	}
    }

*/
import "C"
func main() {
	C.hello()
}