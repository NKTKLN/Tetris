build:
	rm -rf bin
	mkdir bin
	GOOS=windows GOARCH=amd64 go build -o bin/tetris-amd64-win.exe cmd/tetris/main.go 
	GOOS=darwin GOARCH=amd64 go build -o bin/tetris-amd64-macos cmd/tetris/main.go 
	GOOS=linux GOARCH=amd64 go build -o bin/tetris-amd64-linux cmd/tetris/main.go 

run:
	go run cmd/tetris/main.go 
