build:
	go build -o ptree

run: build
	./ptree

clean: ptree
	rm -r ptree