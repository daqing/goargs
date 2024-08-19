run:
  go build -o bin/
  find . -name '*.go' | ./bin/goargs wc -l
  find . -name '*.go' | ./bin/goargs wc -l :1
  find . -name '*.go' | awk -F/ '{print $1, $2}' | ./bin/goargs echo :1/:2
