// main.go

package main

import (
	"fmt"
	"kazokku-app/config"
	"kazokku-app/app"
	"kazokku-app/database"
	"kazokku-app/app/handlers" // Import handlers package
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	// Inisialisasi aplikasi Echo
	e := app.Init()

	// Koneksi ke database
	err = database.Init(cfg)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// Tambahkan route untuk registrasi pengguna
	e.POST("/user/register/", handlers.UserRegisterHandler)
	e.GET("/users", handlers.GetAllUser)


	// Mulai server Echo
	e.Start(fmt.Sprintf(":%d", cfg.ServerPort))
}
