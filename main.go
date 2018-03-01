package main

import (
	"net/http"
	"os"
	"time"

	"github.com/nlopes/slack"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/juntaki/firestarter-sqs-proxy/lib"
	"github.com/juntaki/firestarter/application"
	"github.com/juntaki/firestarter/domain"
	"github.com/juntaki/firestarter/infrastructure"
	proto "github.com/juntaki/firestarter/proto"
)

func main() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic("logger initialize failed")
	}
	logger := zapLogger.Sugar()

	token := os.Getenv("SLACK_TOKEN")
	if len(token) == 0 {
		logger.Error("SLACK_TOKEN is required")
	}

	verificationToken := os.Getenv("SLACK_VERIFICATION_TOKEN")
	if len(verificationToken) == 0 {
		logger.Error("SLACK_VERIFICATION_TOKEN is required")
	}

	// SQS mode enabled?
	var proxy *lib.SQSProxy
	sqsMode := false
	url := os.Getenv("SQS_URL")
	if len(url) != 0 {
		proxy, err = lib.NewSQSProxy(url, "http://localhost:3000")
		if err == nil {
			sqsMode = true
		}
	}

	slackAPI := slack.New(token)
	configRepository := infrastructure.NewConfigRepositoryImpl()
	chatRepository := &infrastructure.ChatRepositorySlackImpl{API: slackAPI}

	bot := application.NewSlackBot(
		verificationToken,
		slackAPI,
		configRepository,
		logger,
		sqsMode,
	)

	botRouter := chi.NewRouter()
	botRouter.Use(middleware.RequestID)
	botRouter.Use(middleware.RealIP)
	botRouter.Use(middleware.Logger)
	botRouter.Use(middleware.Recoverer)
	botRouter.Use(middleware.Timeout(60 * time.Second))

	adminRouter := chi.NewRouter()
	adminRouter.Use(middleware.RequestID)
	adminRouter.Use(middleware.RealIP)
	adminRouter.Use(middleware.Logger)
	adminRouter.Use(middleware.Recoverer)
	adminRouter.Use(middleware.Timeout(60 * time.Second))

	// Interarcitve message API, Slack <-> bot
	botRouter.Post("/", bot.InteractiveMessageHandler)

	// admin API, admin <-> firestarter
	apiHandler := proto.NewConfigServiceServer(&application.AdminAPI{
		ConfigRepository: configRepository,
		ChatRepository:   chatRepository,
		Validator:        domain.NewValidator(),
	}, nil)
	adminRouter.Mount("/twirp/", apiHandler)
	adminRouter.Mount("/", http.FileServer(http.Dir("admin/dist")))

	eg := errgroup.Group{}
	// start RTM event checker
	eg.Go(func() error { return bot.Run() })
	// start HTTP server for interactive message.
	eg.Go(func() error { return http.ListenAndServe(":3000", botRouter) })
	// start HTTP server for admin.
	eg.Go(func() error { return http.ListenAndServe(":8080", adminRouter) })
	// start SQS proxy, if enabled
	if sqsMode {
		eg.Go(func() error { return proxy.Run() })
	}

	err = eg.Wait()
	if err != nil {
		logger.Fatalw("application", zap.Error(err))
	}
}
