build:
	go build -o bin/cgemini .

test:
	go test ./...

run:
	go run .

clean:
	rm -f bin/cgemini
