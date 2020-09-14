clean:
	@ go clean -testcache

run: build
	@ ./bin/hello-fresh

build:
	@ mkdir -p ./bin
	@ go build -o bin/hello-fresh
	@ echo "Binary generated at bin/hello-fresh"