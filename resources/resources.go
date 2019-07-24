package resources

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

func DatabaseConnection() *sql.DB {
	viper.SetConfigName("postgres-dev")
	viper.AddConfigPath("environment/postgres")

	handleReadingConfigFile := viper.ReadInConfig()

	if handleReadingConfigFile != nil {
		log.Printf("Error when read files %s", handleReadingConfigFile)
	}

	var PostgresHost = viper.GetString("postgres.PostgresHost")
	var PostgresUser = viper.GetString("postgres.PostgresUsername")
	var PostgresPass = viper.GetString("postgres.PostgresPassword")
	var PostgresDB = viper.GetString("postgres.PostgresDatabaseName")

	var databaseConfig = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		PostgresUser, PostgresPass, PostgresHost, PostgresDB)

	databaseConnectionConfiguration, errorDatabaseConfiguration := sql.Open("postgres", databaseConfig)

	if errorDatabaseConfiguration != nil {
		log.Printf("Error when connecting DB %s", errorDatabaseConfiguration)
	}

	databaseConnectionConfiguration.SetMaxOpenConns(50)
	databaseConnectionConfiguration.SetMaxIdleConns(50)
	databaseConnectionConfiguration.SetConnMaxLifetime(100 * time.Millisecond)

	return databaseConnectionConfiguration
}
