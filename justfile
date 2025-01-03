alias t:=test
alias b:=build

test:
    go test rwf
    go test rwf/pkg/raider
build:
    go build -o build/test/rwf rwf
release:
    go build -a -gcflags=all="-l -B" -ldflags="-w -s" -o build/release/rwf rwf