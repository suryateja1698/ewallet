
run:  ## Run go server
	go run server.go

generate:  ## generate gql
	go get github.com/99designs/gqlgen@v0.17.16
	go get github.com/99designs/gqlgen/codegen@v0.17.16
	go get github.com/99designs/gqlgen/codegen@v0.17.16
	go get github.com/99designs/gqlgen/internal/imports@v0.17.16
	go get github.com/99designs/gqlgen/codegen/config@v0.17.16
	go get github.com/99designs/gqlgen/internal/imports@v0.17.16
	go run github.com/99designs/gqlgen generate
