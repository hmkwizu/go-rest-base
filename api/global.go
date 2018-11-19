package api

import (
	"database/sql"
	"log"

	// mysql is initialized in InitDB function
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// Configuration holds configuration variables
type Configuration struct {
	//Port to listen to
	Port string

	//ConnectionString to the database
	ConnectionString string

	//Driver for database/sql
	Driver string

	//SignKey for generating JWT Tokens
	SignKey []byte
}

// Config holds configuration variables
var Config Configuration

// db, the database connection
var db *sql.DB

// InitDB opens the database connection, calls InitConfig if necessary
func InitDB() {

	//make sure InitConfig was called
	if Config.ConnectionString == "" || Config.Driver == "" {
		InitConfig()
	}

	var err error
	db, err = sql.Open(Config.Driver, Config.ConnectionString)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

// InitConfig initializes configuration variables
func InitConfig() {

	viper.AutomaticEnv()
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("CONNECTION_STRING", "root@tcp(127.0.0.1:3306)/niru?charset=utf8")
	viper.SetDefault("DRIVER", "mysql")
	viper.SetDefault("SIGN_KEY", "123456789")

	Config.Port = viper.GetString("PORT")
	log.Println("Server is running on PORT " + Config.Port)

	Config.ConnectionString = viper.GetString("CONNECTION_STRING")
	Config.Driver = viper.GetString("DRIVER")
	Config.SignKey = []byte(viper.GetString("SIGN_KEY"))

}
