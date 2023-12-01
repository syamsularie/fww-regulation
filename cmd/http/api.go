package main

import (
	"fmt"
	"fww-regulation/config"
	"fww-regulation/config/middleware"
	"fww-regulation/internal/handler"
	"fww-regulation/internal/repository"
	"fww-regulation/internal/usecase"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	baseDep := config.NewBaseDep()
	loadEnv(baseDep.Logger)
	db, err := config.NewDbPool(baseDep.Logger)

	if err != nil {
		os.Exit(1)
	}

	dbCollector := middleware.NewStatsCollector("fww", db)
	prometheus.MustRegister(dbCollector)
	fiberProm := middleware.NewWithRegistry(prometheus.DefaultRegisterer, "fww-core", "", "", map[string]string{})

	//=== repository lists start ===//
	blacklistRepo := repository.NewBlacklistRepository(repository.BlacklistRepository{
		DB: db,
	})
	//=== repository lists end ===//

	//=== usecase lists start ===//
	dukcapilUsecase := usecase.NewDukcapilUsecase(&usecase.DukcapilUsecase{})

	blacklistUsecase := usecase.NewBlacklistUsecase(&usecase.BlacklistUsecase{
		BlacklistRepo: blacklistRepo,
	})

	peduliLindungiUsecase := usecase.NewPeduliLindungiUsecase(&usecase.PeduliLindungiUsecase{})
	//=== usecase lists end ===//

	//=== handler lists start ===//
	dukcapilHandler := handler.NewDukcapilHandler(handler.Dukcapil{
		DukcapilUsecase: dukcapilUsecase,
	})

	blacklistHandler := handler.NewBlacklistHandler(handler.Blacklist{
		BlacklistUsecase: blacklistUsecase,
	})

	peduliLindungiHandler := handler.NewPeduliLindungiHandler(handler.PeduliLindungi{
		PeduliLindungiUsecase: peduliLindungiUsecase,
	})
	//=== handler lists end ===//
	app := fiber.New(fiber.Config{
		BodyLimit: 30 * 1024 * 1024,
	})

	app.Use(fiberProm.Middleware)
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(pprof.New())
	app.Use(logger.New(logger.Config{
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeInterval: time.Millisecond,
		TimeFormat:   "02-01-2006 15:04:05",
		TimeZone:     "Indonesia/Jakarta",
	}))
	// Define a route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	//=== healthz route
	// app.Get("/", Healthz)
	app.Get("/healthz", Healthz)

	//Dukcapil Routes
	app.Get("/check/dukcapil", dukcapilHandler.CheckDukcapilByKTP)

	//Blacklist Routes
	app.Get("/check/blacklist", blacklistHandler.CheckBlacklist)

	//Peduli Lindungi Routes
	app.Get("/check/pedulilindungi", peduliLindungiHandler.CheckPeduliLindungi)

	//=== listen port ===//
	if err := app.Listen(fmt.Sprintf(":%s", "3000")); err != nil {
		log.Fatal(err)
	}

}

func Healthz(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Service is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}

func loadEnv(logger config.Logger) {
	_, err := os.Stat(".env")
	if err == nil {
		err = godotenv.Load()
		if err != nil {
			logger.Error("no .env files provided")
		}
	}
}
