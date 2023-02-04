install:
	cd ./cmd/myanalyzer && go install

do: install
	go vet -vettool=`which myanalyzer` ./...

fmt:
	go fmt ./...

test: fmt
	go test -v ./...
