all:
	gofmt -e -s -w .
	$(CC) -g -fPIC -c -o lib/lib.o lib/lib.c
	$(CC) -g -fPIC -shared -o liblib.so lib/lib.o
	go vet .
	LD_LIBRARY_PATH=. go run test.go

clean:
	rm -rf lib/*.o *.so
