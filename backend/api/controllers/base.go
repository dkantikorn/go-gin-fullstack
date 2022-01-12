package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dkantikorn/go-gin-fullstack/api/middlewares"
	"github.com/dkantikorn/go-gin-fullstack/api/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

var errList = make(map[string]string)

//Function mking for automatic create database
func (server *Server) CreateDatabase(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) error {
	if strings.ToLower(Dbdriver) == "mysql" {
		connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
			DbUser,
			DbPassword,
			DbHost,
			DbPort)

		db, err := gorm.Open(Dbdriver, connStr)
		_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + DbName + ";")

		if err != nil {
			return err
		}
		return nil
	} else if strings.ToLower(Dbdriver) == "postgres" {
		connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			DbUser,
			DbPassword,
			DbHost,
			DbPort,
			"postgres")

		// connect to the postgres db just to be able to run the create db statement
		db, err := gorm.Open(Dbdriver, connStr)
		if err != nil {
			return err
		}

		// close db connection
		defer db.DB().Close()
		// check if db exists
		type Result struct {
			CountResult int
		}

		var result Result
		rs := db.Debug().Raw("SELECT count(datname) AS count_result FROM pg_database WHERE datname = ?;", DbName).Scan(&result)
		if rs.Error != nil {
			return rs.Error
		}

		//Check for if database if not exists create for the database
		if result.CountResult == 0 {
			stmt := fmt.Sprintf("CREATE DATABASE %s WITH OWNER %s ENCODING='UTF8'", DbName, DbUser)
			rs := db.Debug().Exec(stmt)
			if rs.Error != nil {
				return rs.Error
			}
		}
		return nil
	}
	return nil
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	// If you are using mysql, i added support for you here(dont forgot to edit the .env file)
	if Dbdriver == "mysql" {
		if err := server.CreateDatabase(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName); err != nil {
			fmt.Printf("Cannot create database %s with %s database server", DbName, Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			var err error
			DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
			server.DB, err = gorm.Open(Dbdriver, DBURL)
			if err != nil {
				fmt.Printf("Cannot connect to %s database", Dbdriver)
				log.Fatal("This is the error:", err)
			} else {
				fmt.Printf("We are connected to the %s database", Dbdriver)
			}
		}
	} else if Dbdriver == "postgres" {
		if err := server.CreateDatabase(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName); err != nil {
			fmt.Printf("Cannot create database %s with %s database server", DbName, Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			var err error
			DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
			server.DB, err = gorm.Open(Dbdriver, DBURL)
			if err != nil {
				fmt.Printf("Cannot connect to %s database", Dbdriver)
				log.Fatal("This is the error connecting to postgres:", err)
			} else {
				fmt.Printf("We are connected to the %s database", Dbdriver)
			}
		}
	} else {
		fmt.Println("Unknown Driver")
	}

	//database migration
	server.DB.Debug().AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.ResetPassword{},
		&models.Like{},
		&models.Comment{},
	)

	server.Router = gin.Default()
	server.Router.Use(middlewares.CORSMiddleware())

	server.initializeRoutes()

}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
