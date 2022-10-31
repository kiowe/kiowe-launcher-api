package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"os/signal"
)

type Server struct {
	app  *fiber.App
	addr string
	pool *pgxpool.Pool
}

func NewServer(addr string, app *fiber.App, pool *pgxpool.Pool) *Server {
	return &Server{
		app:  app,
		addr: addr,
		pool: pool,
	}
}

func (s *Server) StartWithGracefulShutdown() {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		log.Println("Closing all database connections...")
		s.pool.Close()
		log.Println("All database connections have been closed.")

		if err := s.app.Shutdown(); err != nil {
			log.Println("[ERR]: The server did not shut down.")
		}

		log.Println("Server has successfully shut down.")

		close(idleConnsClosed)
	}()

	if err := s.app.Listen(s.addr); err != nil {
		log.Printf("[ERR]: Server is not running! Reason: %v\n")
	}

	<-idleConnsClosed
}
