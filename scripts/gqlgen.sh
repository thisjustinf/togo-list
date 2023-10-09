#!/bin/bash
printf "\nRegenerating gqlgen files\n"
rm -f internal/graphql/generated/generated.go \
    internal/graphql/models/models_gen.go
time go run -v github.com/99designs/gqlgen $1
printf "\nDone.\n\n"