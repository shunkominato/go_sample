package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"src/dataloader"
	"src/graph"
	"src/graph/model"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	companyLoader := dataloader.NewCompanyLoader(dataloader.CompanyLoaderConfig{
		MaxBatch: 100, // 溜める最大数、0を指定すると制限無し
		Wait:     2 * time.Millisecond, // 溜める時間
		Fetch: func(keys []string) ([]*model.Company, []error) {
				companies := make([]*model.Company, len(keys))
				errors := make([]error, len(keys))

				// 取得処理を書く SELECT * FROM company WHERE company_id IN (...)

				// 引数のkeysに対応する順番の配列で返す。
				return companies, errors
		},
})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		err.Message = e.Error()
		err.Extensions = map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
		}
		return err
	})
	
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
