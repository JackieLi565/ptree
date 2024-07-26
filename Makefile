build:
	go build -o ptree

run: build
	./ptree

tesst:
	go test

clean: ptree
	rm -r ptree