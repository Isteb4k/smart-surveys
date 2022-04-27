.PHONY: server client

M = $(shell printf "\033[34;1m▶\033[0m")

server:
	go run ./cmd/main.go

gqlgen: ; $(info $(M) Starting GraphQL schemas generation...)
	make gqlgenModerator
	make gqlgenRespondent

gqlgenModerator: ; $(info $(M) Generate for moderator...)
	(cd transport/graphql/moderator && go get github.com/99designs/gqlgen && go run github.com/99designs/gqlgen generate)

gqlgenRespondent: ; $(info $(M) Generate for respondent...)
	(cd transport/graphql/respondent && go get github.com/99designs/gqlgen && go run github.com/99designs/gqlgen generate)