package main

import (
	"flag"
	"forum/global"
	"forum/internal/model"
	"forum/internal/routers"
	"forum/pkg/logger"
	"forum/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"time"
)

var configPath = flag.String("config", "", "Config file path")

// @title forum
// @version 0.01
func main() {
	flag.Parse()

	if *configPath == "" {
		*configPath = "./config/config-dev.yaml"
	}

	appInit(*configPath)

	router := routers.NewRouter()
	s := http.Server{
		Addr:    ":" + global.ServerSetting.HttpPort,
		Handler: router,
	}
	log.Fatal(s.ListenAndServe())
}

func appInit(path string) {
	err := setupSetting(path)
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDbEngine err: %v", err)
	}
}

func setupSetting(path string) error {
	s, err := setting.NewSetting(path)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	httpPort := os.Getenv("HttpPort")
	if len(httpPort) != 0 {
		global.ServerSetting.HttpPort = httpPort
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost != "" {
		global.DatabaseSetting.Host = dbHost
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword != "" {
		global.DatabaseSetting.Password = dbPassword
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret != "" {
		global.JWTSetting.Secret = jwtSecret
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	return err
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
