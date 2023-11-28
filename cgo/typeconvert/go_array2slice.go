package main

/*
#include <stdio.h>
#include <stdlib.h>

struct array {
	int i;
	int j;
};
int numbers[5] = {11,22,33,44,55};
struct array* GetArray(int n) {
	struct array *ptr =  malloc(n* sizeof(struct array));
	for(int i = 0; i < n; i++) {
		(ptr + i)->i = i;
		(ptr + i)->j = i;
	}
	return ptr;

	for (int i = 0; i < 5; i++) {
		numbers[i] = i + 10;
	}
}
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

type array struct {
	i int
	j int
}

func main() {
	var array *C.struct_array = C.GetArray(C.int(4))
	unSafePtr := unsafe.Pointer(array)
	defer C.free(unSafePtr)

	arrayPtr := (*[1 << 30]C.struct_array)(unSafePtr)
	slice := arrayPtr[0:4:4]
	fmt.Println(slice)

	var arr0 []int
	arr0Header := (*reflect.SliceHeader)(unsafe.Pointer(&arr0))
	arr0Header.Data = uintptr(unsafe.Pointer(&C.numbers[0]))
	arr0Header.Len = 5
	arr0Header.Cap = 5
	fmt.Println(C.numbers)
	fmt.Println(arr0Header)
	fmt.Println(arr0)

	numArray := (*[1 << 30]C.int)(unsafe.Pointer(&C.numbers[0]))
	numSlice := numArray[0:5:5]
	fmt.Println(numSlice)
}
