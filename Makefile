run:
	#go run cmd/web/* -addr=":8081"
	go run `ls cmd/web/*.go | grep -v _test.go` -addr=":8081"



#go run $(ls -1 cmd/web/*.go -addr=":8081" | grep -v _test.go) -addr=":8081"