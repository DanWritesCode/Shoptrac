package main

import (
	"./config"
	"./database"
	"./logging"
	"./shopify"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime"
	"strconv"
)

// DB is the exported DB value
var log *logrus.Logger

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	log = logging.GetLogger()

	cfg, err := config.NewConfig("config.json")
	if err != nil {
		log.WithError(err).Fatalf("Error getting config from path %v.", "config.json")
		return
	}

	shopify.SetupClient()
	shopify.NewShopify(cfg.App.ClientId, cfg.App.ClientSecret, cfg.App.RedirectURL)

	err = logging.SetLogPath(cfg.LogFile)
	if err != nil {
		log.WithError(err).Fatalf("Error settings log path.")
		return
	}

	if cfg.Debug {
		logging.SetLevel(logging.DebugLevel)
	}

	database.ConnectToDB(cfg.DB.Username, cfg.DB.Password, cfg.DB.Address, cfg.DB.Database)
	defer database.CloseDB()

	shopify.TestDataImports()

	//limit = limiter.NewConcurrencyLimiter(2000)

	r := InitRouter(cfg.Timeout)

	log.Infof("Now listening on port %v.", cfg.Port)
	fmt.Println(fmt.Sprintf("Now listening on port %v.", cfg.Port))
	http.ListenAndServe(cfg.Host+":"+strconv.Itoa(cfg.Port), r)
}
