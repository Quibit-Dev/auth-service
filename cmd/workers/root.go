package workers

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"auth-service/internal/config"
	"auth-service/internal/pkg/logger"
	"auth-service/internal/provider"

	"github.com/spf13/cobra"
)

func Start() {
	cfg := config.InitConfig()
	logger.SetZerologLogger(cfg)
	provider.InitGlobal(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	rootCmd := &cobra.Command{}
	cmd := []*cobra.Command{
		{
			Use:   "serve-all",
			Short: "Run all",
			Run: func(cmd *cobra.Command, _ []string) {
				runHttpWorker(cfg, ctx)
			},
		},
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
