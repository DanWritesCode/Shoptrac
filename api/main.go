package main

import (
  "./config"
  "./endpoints"
  "./logging"
  "fmt"
  "github.com/jmoiron/sqlx"
  "github.com/sirupsen/logrus"
  "net/http"
  "runtime"
  "strconv"
  "time"

  _ "github.com/Go-SQL-Driver/MySQL"
  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "github.com/go-chi/cors"
)

// DB is the exported DB value
var DB *sqlx.DB
var log *logrus.Logger
var TblList []string

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())

  log = logging.GetLogger()

  cfg, err := config.NewConfig("config.json")
  if err != nil {
    log.WithError(err).Fatalf("Error getting config from path %v.", "config.json")
    return
  }

  //fireproxies.InitGlobals()

  err = logging.SetLogPath(cfg.LogFile)
  if err != nil {
    log.WithError(err).Fatalf("Error settings log path.")
    return
  }

  if cfg.Debug {
    logging.SetLevel(logging.DebugLevel)
  }

  ConnectToDB(cfg.DB.Username, cfg.DB.Password, cfg.DB.Address, cfg.DB.Database)
  defer CloseDB()

  go TableListTimer()

  //limit = limiter.NewConcurrencyLimiter(2000)

  r := chi.NewRouter()

  cors := cors.New(cors.Options{
    AllowedOrigins:   []string{"*"},
    AllowedMethods:   []string{"GET", "POST"},
    AllowedHeaders:   []string{"Accept", "Content-Type"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: true,
    MaxAge:           300,
  })

  r.Use(
    middleware.Recoverer,
    middleware.Timeout(time.Duration(cfg.Timeout)*time.Second),
    cors.Handler,
  )

  r.Route("/", func(r chi.Router) {
    // Protected routes
    r.Group(func(r chi.Router) {
      // Shopify App Install Endpoint - handles OAuth with Shopify
      r.Post("/install", endpoints.PostInstall)

      // Dashboard GET Endpoint - Shows Summary
      r.Get("/summary", endpoints.GetSummary)

      // Items/Products view page
      r.Get("/items", endpoints.GetItems)

      r.Get("/orders", endpoints.GetOrders)

      r.Get("/customers", endpoints.GetCustomers)

      r.Get("/revenue", endpoints.GetRevenue)

      r.Get("/expenses", endpoints.GetExpenses)

      // Configuration POST - User Edits StonksUp' Configuration
      r.Post("/config", endpoints.PostConfig)

    })
    r.NotFound(endpoints.Error404)
    r.MethodNotAllowed(endpoints.Error405)
  })

  log.Infof("Now listening on port %v.", cfg.Port)
  fmt.Println(fmt.Sprintf("Now listening on port %v.", cfg.Port))
  http.ListenAndServe(cfg.Host+":"+strconv.Itoa(cfg.Port), r)
}

// ConnectToDB will establish a connection to the database.
func ConnectToDB(user, password, address, dbname string) {
  dab, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True", user, password, address, dbname))
  if err != nil {
    log.Fatalln(err)
  }
  dab.SetMaxOpenConns(5000)
  dab.SetMaxIdleConns(400)
  DB = dab
}

func GetTableList() {
  TblList = nil
  _ = DB.Select(&TblList, "SHOW TABLES;")
}

func TableListTimer() {
  GetTableList()
  for range time.Tick(time.Hour *1){
    GetTableList()
  }
}


// CloseDB cleanly closes the DB connection.
func CloseDB() error {
  return DB.Close()
}
