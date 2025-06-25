package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/cors"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"

	"go-auth-api/pkg/config"

	"github.com/go-chi/chi/v5/middleware"
	"go-auth-api/internal/user/handler"
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

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Front React
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // 5 minutos
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	authHandler := handler.NewAuthHandler(db)

	//Teste de rota viva

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🔥 API viva e rodando!"))
	})

	r.Post("/signup", authHandler.SignupHandler)
	r.Post("/login", authHandler.LoginHandler)

	fmt.Println("Servidor rodando na porta", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
