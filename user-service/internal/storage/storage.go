package storage

import (
	"database/sql"
	"log"
	"os"
	"user-service/internal/infrastructura/postgres"
	"user-service/internal/infrastructura/redis"
	"user-service/internal/service"
	userservice "user-service/user_service"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func OpenSql(driverName, url string) (*sql.DB, error) {
	db, err := sql.Open(driverName, url)
	if err != nil {
		log.Println("failed to open database:", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println("Unable to connect to database:", err)
		return nil, err
	}
	return db, err
}

func Handler(db *sql.DB) *userservice.Service {

	redisClient := redis.NewRedisClient(os.Getenv("redis_url"), "", 0)

	repo := postgres.NewUserPostgres(db)
	s := service.NewUserService(repo)
	user_handler := userservice.NewService(s, redisClient)
	return user_handler
}
