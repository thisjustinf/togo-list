#!/bin/bash
printf "\nRegenerating gqlgen files\n"
rm -f internal/graphql/generated/exec.go \
    internal/graphql/models/generated.go \
    internal/graphql/resolvers/generated/generated.go
time go run -v github.com/99designs/gqlgen $1
printf "\nDone.\n\n"