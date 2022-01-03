
dependencies:
	go get

build: dependencies dynbio.go
	go build -o build/dymbio dynbio.go

clean:
	go clean
	rm -rf build
