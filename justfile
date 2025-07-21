entry := "cmd/driftc/main.go"

run *ARGS:
    @go run {{entry}} {{ARGS}}

test:
    @go test pkg/driftc/*.go
