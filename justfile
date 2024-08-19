run:
  go build -o bin/
  find . -name '*.go' | ./bin/goargs wc -l
