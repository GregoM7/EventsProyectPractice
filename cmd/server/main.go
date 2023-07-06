package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/GregoM7/EventsProyectPractice/cmd/server/controller"
	"github.com/GregoM7/EventsProyectPractice/internal/event"
	"github.com/GregoM7/EventsProyectPractice/internal/user"
	"github.com/GregoM7/EventsProyectPractice/package/middleware"
	"github.com/GregoM7/EventsProyectPractice/package/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	var (
		ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("user"),
			os.Getenv("pass"),
			os.Getenv("hostdb"),
			os.Getenv("port"),
			os.Getenv("db_name"))
	)
	fmt.Print(ConnectionString)

	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		log.Fatal("Error opening database")
	}
	
	storeSQL := store.NewSQLStore(db)

	//Users
	repoUsers := user.NewRepository(storeSQL)
	serviceUsers := user.NewService(repoUsers)
	controllerUsers := controller.NewUserController(serviceUsers)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	users := r.Group("/users")
	{
		users.GET("", controllerUsers.ReadAll())
		users.POST("",middleware.AuthenticationMiddleware(), controllerUsers.Create())
	}

	//Events
	repoEvents := event.NewRepository(storeSQL)
	serviceEvents := event.NewService(repoEvents)
	controllerEvents := controller.NewEventController(serviceEvents, serviceUsers)

	events := r.Group("/events")
	{
		events.GET("",controllerEvents.ReadAll())
	}
	r.Run(":8080")
}