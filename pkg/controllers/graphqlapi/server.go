package graphqlapi

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sede-x/gopoc-connector/pkg/controllers/graphqlapi/graph"
	"github.com/sede-x/gopoc-connector/pkg/logic"
)

type Server struct {
	logic.ConnectorLogic
}

func (s *Server) Start(serveraddr string) {
	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{ConnectorLogic: s.ConnectorLogic}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", server)

	// start server
	log.Printf("Connector GraphQL API is running at http://localhost%s/", serveraddr)
	http.ListenAndServe(serveraddr, nil)
}
