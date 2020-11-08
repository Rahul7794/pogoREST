package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "pogoREST/docs" // Required for side effects
	"pogoREST/errors"
	"pogoREST/handlers"
)

// Command line Args
var runCmd = &cobra.Command{
	Use:   "serve",                        // SubCommand
	Short: "rest interface over postgres", // Short description of the SubCommand
	Long:  "rest interface over postgres which supports various DB Operations ",
	RunE:  pogoREST,
}

// dbType is the engine type such as postgres, mysql, sql
var dbType string

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&dbType, "db", "d", "", "database type to connect")
	err := runCmd.MarkFlagRequired("db")
	handleCommandError(err)
}

// @title User Application
// @description This is a user management application
// @version 1.0
// @host localhost:8081
// @BasePath /api/v1
// pogoREST is rest interface over postgres with swagger support for UI
func pogoREST(_ *cobra.Command, _ []string) error {
	e := echo.New()
	e.HTTPErrorHandler = errors.ErrorHandler
	// Base path for the API
	v1 := e.Group("/api/v1")
	{
		// Creates an object of handler class
		db, err := handlers.NewDBEngine(dbType)
		if err != nil {
			return err
		}
		v1.POST("/user", db.CreateUser)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// starting the service
	e.Logger.Fatal(e.Start(":8081"))
	return nil
}
