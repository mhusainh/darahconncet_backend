package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/mhusainh/DarahConnect/DarahConnectAPI/configs"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/builder"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/service"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/cache"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/cloudinary"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/database"
	googleoauth "github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/googleOauth"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/mailer"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/midtrans"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/server"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/timezone"
)

func main() {
	cfg, err := configs.NewConfig(".env")
	// cfg, err := configs.NewConfigYaml("config.yaml")
	checkError(err)

	db, err := database.InitDatabase(cfg.PostgresConfig)
	checkError(err)

	rdb := cache.InitCache(cfg.RedisConfig)

	err = timezone.InitTimezone()
	checkError(err)

	cloudinaryService, err := cloudinary.NewService(&cfg.CloudinaryConfig)
	checkError(err)

	mailer, err := mailer.NewMailer(&cfg.SMTPConfig)
	checkError(err)

	err = googleoauth.InitGoogle(&cfg.GoogleOauth)
	checkError(err)

	_, err = midtrans.InitMidtrans(&cfg.MidtransConfig)
	checkError(err)

	blockchain, err := service.NewBlockchainService(cfg.Blockchain)
	checkError(err)

	publicRoutes := builder.BuildPublicRoutes(cfg, db, rdb, cloudinaryService, mailer, blockchain)
	privateRoutes := builder.BuildPrivateRoutes(cfg, db, rdb, cloudinaryService, mailer, blockchain)

	srv := server.NewServer(cfg, publicRoutes, privateRoutes)
	runServer(srv, cfg.PORT)
	waitForShutdown(srv)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func runServer(srv *server.Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		log.Fatal(err)
	}()
}

func waitForShutdown(srv *server.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal(err)
		}
	}()
}
