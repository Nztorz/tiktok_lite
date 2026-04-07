/cmd/
  /api/
    main.go

/internal/
  /handlers/      → HTTP layer
  /services/      → business logic
  /repository/    → DB access (uses sqlc)
  /models/        → shared structs
  /middleware/    → auth, logging, etc.
  /auth/          → JWT, password, tokens
  /storage/       → S3 logic
  /video/         → ffmpeg processing

/sql/
  /schema/        → tables
  /queries/       → SQL queries (for sqlc)

/migrations/      → DB migrations (goose)

go.mod

FLOW of request:

Client (request)
      ↓
Handler
      ↓
Service
      ↓
Repository → Database
      ↓
Storage (S3) / Video (ffmpeg)
      ↓
Service
      ↓
Handler
      ↓
Client (response)

users.sql

ID uuid primary key default gen_random_uuid() 
email text not null unique
bio text
hashed_password TEXT not null
created at timestamptz not null default now()
updated at timestamptz not null default now()


POST /api/public/user
- Create a new sqlc query to create a new user
- Create model for user request
- Read the request and decode in struct
- create new user in db
- response




func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Auth check")
		next.ServeHTTP(w, r)
	})
}

func ProtectedRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("secret"))
	})

	return AuthMiddleware(mux)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/", ProtectedRoutes())

	http.ListenAndServe(":8080", mux)
}