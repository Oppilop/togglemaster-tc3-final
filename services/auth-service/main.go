package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
<<<<<<< HEAD
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // registra o driver pgx
	"github.com/joho/godotenv"
=======
	"github.com/joho/godotenv"
	_ "github.com/jackc/pgx/v5/stdlib" // registra o driver pgx
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
)

// App struct (para injeção de dependência)
type App struct {
<<<<<<< HEAD
	DB        *sql.DB
	MasterKey string
=======
	DB         *sql.DB
	MasterKey  string
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
}

func main() {
	// Carrega o .env para desenvolvimento local. Em produção, isso não fará nada.
	_ = godotenv.Load()

	// --- Configuração ---
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" // Porta padrão
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL deve ser definida")
	}

	masterKey := os.Getenv("MASTER_KEY")
	if masterKey == "" {
		log.Fatal("MASTER_KEY deve ser definida")
	}

	// --- Conexão com o Banco ---
	db, err := connectDB(databaseURL)
	if err != nil {
		log.Fatalf("Não foi possível conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	app := &App{
<<<<<<< HEAD
		DB:        db,
		MasterKey: masterKey,
=======
		DB:         db,
		MasterKey:  masterKey,
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
	}

	// --- Rotas da API ---
	mux := http.NewServeMux()
	mux.HandleFunc("/health", app.healthHandler)

	// Endpoint público para validar uma chave
	mux.HandleFunc("/validate", app.validateKeyHandler)

<<<<<<< HEAD
	// Endpoints de "admin" protegidos pelo middleware MasterKey
	mux.Handle("/admin/keys", app.masterKeyAuthMiddleware(http.HandlerFunc(app.createKeyHandler)))

	// SECURITY (gosec G114): http.ListenAndServe sem timeouts deixa o servidor
	// vulnerável a Slowloris. Construímos explicitamente um http.Server com
	// ReadHeaderTimeout, ReadTimeout, WriteTimeout e IdleTimeout.
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	log.Printf("Serviço de Autenticação (Go) rodando na porta %s", port)
	if err := server.ListenAndServe(); err != nil {
=======
	// Endpoints de "admin" para criar/gerenciar chaves
	// Eles são protegidos pelo middleware de autenticação
	mux.Handle("/admin/keys", app.masterKeyAuthMiddleware(http.HandlerFunc(app.createKeyHandler)))

	log.Printf("Serviço de Autenticação (Go) rodando na porta %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
>>>>>>> 9204ee7da5de7b4a037661c564a1748b6d514677
		log.Fatal(err)
	}
}

// connectDB inicializa e testa a conexão com o PostgreSQL
func connectDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conectado ao PostgreSQL com sucesso!")
	return db, nil
}
