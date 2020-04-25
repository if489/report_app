package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	_ "github.com/lib/pq"

	"reports_app/handlers"
	"reports_app/logic"
)

func main() {

	db, err := sql.Open("postgres", "host=localhost user=postgres password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("./public"))

	r.Get(
		"/",
		fs.ServeHTTP,
	)

	loader := logic.NewReportsLoader(db)
	getReportsLoader := handlers.NewGetReportsLoader(loader)
	r.Get("/reports",
		getReportsLoader.ServeHTTP,
	)

	updater := logic.NewReportUpdater(db)
	reportUpdater := handlers.NewPostPutReportUpdater(updater)
	r.Post(
		"/reports/block/{reportId}",
		reportUpdater.PostServeHTTP,
	)
	r.Put(
		"/reports/{reportId}",
		reportUpdater.PutServeHTTP,
	)

	http.ListenAndServe(":3000", r)
}
