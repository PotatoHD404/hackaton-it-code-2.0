package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/uptrace/bun/dbfixture"
	"hackaton-it-code-2.0/src/handler"
	"os"
)

func main() {
	handler := function.NewAPIHandler()
	books := make([]function.Book, 0)
	ctx := context.Background()
	handler.Db.RegisterModel((*function.Book)(nil), (*function.Author)(nil))
	fixture := dbfixture.New(handler.Db, dbfixture.WithRecreateTables())

	// Load fixture.yaml which contains data for User and Story models.
	if err := fixture.Load(ctx, os.DirFS("."), "fixture.yaml"); err != nil {
		panic(err)
		//log.Fatal(http.ListenAndServe(":3000", handler))
	}

	if err := handler.Db.NewSelect().Model(&books).OrderExpr("id ASC").Scan(ctx); err != nil {
		panic(err)
	}
	//println(ctx)
	res2B, _ := json.Marshal(books)
	fmt.Println(string(res2B))
	//println(books)
}
