//go:build go1.19

package tools

import (
	_ "github.com/daixiang0/gci/cmd/gci"
	_ "github.com/go-task/task/v3/cmd/task"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)

//go:generate go build -v -o ../bin/ github.com/go-task/task/v3/cmd/task
//go:generate go build -v -o ../bin/ github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go build -v -o ../bin/ github.com/daixiang0/gci
