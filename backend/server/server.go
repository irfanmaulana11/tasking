package server

import (
	dbr "be-tasking/app/repository/db"
	"be-tasking/app/service"
	"be-tasking/config"
	"log"

	"net/url"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RESTServer struct {
	router *gin.Engine
}

type RestRunner interface {
	Run()
}

func (rs *RESTServer) Run() {
	servicePort := os.Getenv("SERVICE_PORT")
	log.Printf("Running Server in :%s", servicePort)

	if err := rs.router.Run(":" + servicePort); err != nil {
		log.Fatal(err)
	}
}

func NewRestServer() RestRunner {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware)

	dbPort, _ := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 32)

	mySqlConn, err := dbr.NewMySQLConn(config.MySQLConfiguration{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     int(dbPort),
		DBName:     os.Getenv("DB_NAME"),
		DBOptions:  os.Getenv("DB_OPTIONS"),
		Locale:     url.QueryEscape(os.Getenv("DB_LOCALE")),
	})

	if err != nil {
		log.Println(err)
	}

	// init mysql repo
	mysqlRepo := dbr.NewMySQLRepo(mySqlConn)

	// init services
	hcs := service.NewHealthCheckService()
	aus := service.NewAuthService(mysqlRepo)
	tas := service.NewTaskService(mysqlRepo)

	// init route
	InitRoutes(r, hcs, aus, tas)

	return &RESTServer{
		router: r,
	}
}
