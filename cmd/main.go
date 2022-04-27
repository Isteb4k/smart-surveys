package main

import (
	//_ "github.com/99designs/gqlgen"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"smart-surveys/storage"
	mgql "smart-surveys/transport/graphql/moderator"
	mgen "smart-surveys/transport/graphql/moderator/generated"
	rgql "smart-surveys/transport/graphql/respondent"
	rgen "smart-surveys/transport/graphql/respondent/generated"
	"smart-surveys/usecases/moderator"
	"smart-surveys/usecases/respondent"
)

func main() {
	mux := http.NewServeMux()

	store := storage.NewClient()

	mUseCases := moderator.NewUseCases(store)
	rUseCases := respondent.NewUseCases(store)

	moderatorSrv := handler.NewDefaultServer(mgen.NewExecutableSchema(mgen.Config{Resolvers: &mgql.Resolver{
		UseCases: mUseCases,
	}}))
	respondentSrv := handler.NewDefaultServer(rgen.NewExecutableSchema(rgen.Config{Resolvers: &rgql.Resolver{
		UseCases: rUseCases,
	}}))

	mux.Handle("/moderator/", playground.Handler("GraphQL playground", "/moderator/query"))
	mux.Handle("/respondent/", playground.Handler("GraphQL playground", "/respondent/query"))
	mux.Handle("/moderator/query", moderatorSrv)
	mux.Handle("/respondent/query", respondentSrv)

	log.Printf("connect to http://localhost:%d/respondent for GraphQL playground", 8080)
	log.Printf("connect to http://localhost:%d/moderator for GraphQL playground", 8080)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
