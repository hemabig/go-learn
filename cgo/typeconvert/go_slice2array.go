package main

/*
#include <stdio.h>

void array(int *s, int len) {
	for (int i = 0; i < len; i++) {
		printf("%d-----%d\n", i, *(s+i));
	}
}
*/
import "C"

func main() {
	s := []C.int{5, 6, 7, 8}
	C.array(&s[0], C.int(len(s)))
}
