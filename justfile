driftc_entry := "cmd/driftc/main.go"
driftls_entry := "cmd/driftls/main.go"

run *ARGS:
    @go run {{driftc_entry}} {{ARGS}}

test:
    @go test pkg/driftc/*.go

build-ls:
    @go build -o dist/driftls {{driftls_entry}}
