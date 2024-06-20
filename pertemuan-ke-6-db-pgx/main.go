package main

import (
	"context"
	"fmt"
	"log"
	"praisindo/entity"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// func main() {
// 	gin.SetMode(gin.ReleaseMode)
// 	r := gin.Default()

// 	// // setup service
// 	var mockUserDBInSlice []entity.User
// 	userRepo := slice.NewUserRepository(mockUserDBInSlice)
// 	userService := service.NewUserService(userRepo)
// 	userHandler := handler.NewUserHandler(userService)

// 	// // Routes
// 	router.SetupRouter(r, userHandler)

// 	// Run the server
// 	log.Println("Running server on port 8181")
// 	r.Run(":8181")
// }

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	dsn := "postgresql://postgres:12345678@localhost:5432/Golang"
	pool, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		log.Fatalln(err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("successfully connect to db")

	var u entity.User
	ctx := context.Background()
	// err = pool.QueryRow(ctx, "select id,name from users order by id desc limit 1").Scan(&u.ID, &u.Name)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("user retrieved", u)

	_, err = pool.Exec(ctx, "insert into users (name,email,password,created_at,updated_at) values "+
		"('test3','test3@test.com','pass3',NOW(),NOW())")
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(res.RowsAffected())
	err = pool.QueryRow(ctx, "select id,name,email from users order by id desc limit 1").Scan(&u.ID, &u.Name, &u.Email)
	fmt.Println("user retrieved", u)
	if err != nil {
		log.Fatalln(err)
	}

	// query untuk mengambil row
	rows, err := pool.Query(ctx, "select id,name from users order by id asc")
	var users []entity.User
	for rows.Next() {
		var u2 entity.User
		rows.Scan(&u2.ID, &u2.Name)
		if err != nil {
			log.Println(err)
		}
		users = append(users, u2)
	}
	
	for _, user := range users {
		fmt.Println("name", user.Name, "id", user.ID)
	}
	// fmt.Println("user retrieved", u)
	// // pgxPool, err := connectDB("postgresql://postgres:postgres@localhost:5432/postgres")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// // setup service

	// slice db is disabled. uncomment to enabled
	// var mockUserDBInSlice []entity.User
	// _ = slice.NewUserRepository(mockUserDBInSlice)

	// // pgx db is enabled. comment to disabled
	// userRepo := postgres_pgx.NewUserRepository(pgxPool)
	// userService := service.NewUserService(userRepo)
	// userHandler := handler.NewUserHandler(userService)

	// // Routes
	// router.SetupRouter(r, userHandler)

	// Run the server
	log.Println("Running server on port 8181")
	r.Run(":8181")
}

// func connectDB(dbURL string) (postgres_pgx.PgxPoolIface, error) {
// 	return pgxpool.New(context.Background(), dbURL)
// }
