package main

import (
	"github.com/balaweblog/gosqlapi/handlers"
	//"github.com/balaweblog/gosqlapi/model"
	"log"
	"net/http"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:20UserId20$$pw!@tcp(127.0.0.1:9000)/gotest?charset=utf8")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.Handle("/showalldatabases", handlers.Showalldatabases(db))
	mux.Handle("/createtable", handlers.Createtable(db))
	mux.Handle("/createdb", handlers.Createdb(db))
	mux.Handle("/selecttable", handlers.Selecttable(db))
	mux.Handle("/inserttable", handlers.Inserttable(db))
	mux.Handle("/updatetable", handlers.Updatetable(db))
	mux.Handle("/currentdb", handlers.CurrentInUseDB(db))
	mux.Handle("/showalltables", handlers.ShowalltablesinDb(db))
	mux.Handle("/deletetable", handlers.Deletetable(db))
	mux.Handle("/truncatetable", handlers.Truncatetable(db))
	mux.Handle("/describetable", handlers.Describetableindb(db))
	mux.Handle("/explaintable", handlers.ExplaintableinDb(db))
	mux.Handle("/altertable", handlers.Altertable(db))
	mux.Handle("/showalltableindex", handlers.Showalltableindex(db))
	mux.Handle("/droptable", handlers.Droptable(db))
	mux.Handle("/dropdatabase", handlers.Dropdatabase(db))
	mux.Handle("/listallusers", handlers.Listallusers(db))

	log.Printf("serving on port 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
