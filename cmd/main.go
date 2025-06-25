package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"

	"github.com/krittz/go-auth-api/pkg/config"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Erro ao conectar no banco: ", err)
	}
	defer db.Close()

	r := chi.NewRouter()

	//Teste de rota viva

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🔥 API viva e rodando!"))
	})

	fmt.Println("Servidor rodando na porta", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
