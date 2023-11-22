package cgo_helper

/* #include <stdio.h>
void print(char * s) {
	printf("%s\n", s);
}
*/
import "C"

type CChar C.char

func SayHello(s *C.char) {
	C.print(s)
}
