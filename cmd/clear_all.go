package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"

    "github.com/mikaponics/mikaponics-thing/internal/models"
)

func init() {
    rootCmd.AddCommand(clearAllCmd)
}

var clearAllCmd = &cobra.Command{
    Use:   "clear_all",
    Short: "Clear all data.",
    Long:  `Command drops and recreates all the tables in this web service so the data becomes cleared.`,
    Run: func(cmd *cobra.Command, args []string) {
        // Load up our `environment variables` from our operating system.
        dbHost := os.Getenv("MIKAPONICS_THING_DB_HOST")
        dbPort := os.Getenv("MIKAPONICS_THING_DB_PORT")
        dbUser := os.Getenv("MIKAPONICS_THING_DB_USER")
        dbPassword := os.Getenv("MIKAPONICS_THING_DB_PASSWORD")
        dbName := os.Getenv("MIKAPONICS_THING_DB_NAME")

        // Initialize and connect our database layer for the command.
        dal := models.InitDataAccessLayer(dbHost, dbPort, dbUser, dbPassword, dbName)

        // Drop our tables and recreate them as empty.
        dal.CreateThingTable(true)

        fmt.Println("All data in database where cleared.")
    },
}
