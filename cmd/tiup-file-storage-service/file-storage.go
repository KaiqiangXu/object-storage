package main

import (
	"log"
	"runtime"

	"file-storage-system/adapters"
	"file-storage-system/handlers"
	config "tools/config-parser"
	"tools/version"
)

func main() {
	log.Println("====== STARTING TiUP FILE STORAGE SERVICE ======")

	version.Init()
	cfg, err := config.ParseConfig()
	if err != nil {
		panic(err)
	}

	runtime.GOMAXPROCS(8)
	log.Println("Current GOMAXPROCS:", runtime.GOMAXPROCS(0))

	adapters.SetConfig(cfg.DataAccessLayer.MongoHost)
	handlers.Start(":" + cfg.FileStorage.Port)

	select {}
}
