package main

import (
	"net/http"
	"os"

	"github.com/nlopes/slack"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/juntaki/firestarter/application"
	"github.com/juntaki/firestarter/infrastructure"
	"github.com/juntaki/firestarter/proto"
)

func main() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic("logger")
	}
	logger := zapLogger.Sugar()

	token := os.Getenv("SLACK_TOKEN")
	if len(token) == 0 {
		logger.Panic("token is required")
	}

	verificationToken := os.Getenv("SLACK_VERIFICATION_TOKEN")
	if len(verificationToken) == 0 {
		logger.Panic("verification token is required")
	}

	bot := &application.SlackBot{
		VerificationToken: verificationToken,
		API:               slack.New(token),
		ConfigRepository:  infrastructure.NewConfigRepositoryImpl(),
		Log:               logger,
		Session:           application.NewSession(),
	}

	// Interarcitve message API, Slack <-> bot
	botMux := http.NewServeMux()
	botMux.Handle("/", bot)

	// admin API, admin <-> firestarter
	adminMux := http.NewServeMux()
	apiHandler := proto.NewConfigServiceServer(&application.AdminAPI{
		ConfigRepository: &infrastructure.ConfigRepositoryImpl{},
	}, nil)
	adminMux.Handle("/api/", LogMiddleware(logger, apiHandler))

	adminMux.Handle("/", LogMiddleware(logger, http.FileServer(http.Dir("admin/dist"))))

	eg := errgroup.Group{}
	// start RTM event checker
	eg.Go(func() error { return bot.Run() })
	// start HTTP server for interactive message.
	eg.Go(func() error { return http.ListenAndServe(":3000", botMux) })
	// start HTTP server for admin.
	eg.Go(func() error { return http.ListenAndServe(":8080", adminMux) })

	err = eg.Wait()
	if err != nil {
		logger.Fatalw("application", zap.Error(err))
	}
}

func LogMiddleware(log *zap.SugaredLogger, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infow("Request",
			zap.String("Method", r.Method),
			zap.String("URI", r.RequestURI),
		)
		handler.ServeHTTP(w, r)
	})
}
