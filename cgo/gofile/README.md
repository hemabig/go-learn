go build -o hello.so -buildmode=c-shared hello.go
gcc main.c -o main -lhello