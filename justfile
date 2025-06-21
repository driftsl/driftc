src := "cmd/*.go"

run *ARGS:
    @go run {{src}} {{ARGS}}

test:
    @go test pkg/driftc/*.go
