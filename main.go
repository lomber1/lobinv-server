package main

import (
	"fmt"
	"lobinv-server/api"
	"lobinv-server/config"
	"lobinv-server/database"
	"net/http"
)

func main() {
	c := config.GetConfig()

	database.ConnectDB(c)

	r := api.NewMainRouter()
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", c.Server.Port),
		Handler: r,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
