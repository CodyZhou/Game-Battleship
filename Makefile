# env target is used to build the environment for this Gin framework project.
env:
	go get -u github.com/gin-gonic/gin

# start target is used to build the whole project and run the program.
start:
	go run main.go

# test target is used to run test case.
test:
	go test -v -cover ./test/

.PHONY: env start test
