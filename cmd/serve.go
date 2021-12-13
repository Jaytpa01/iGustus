package cmd

import (
	"github.com/Jaytpa01/iGustus/internal/app"
	"github.com/Jaytpa01/iGustus/internal/config"
	"github.com/Jaytpa01/iGustus/pkg/logger"
	"github.com/Jaytpa01/iGustus/pkg/util"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves the HTTP REST API",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// Verify the environment
		VerifyServeEnvironment()
		err := config.ReadConfig()
		if err != nil {
			logger.Log.Error("error reading config", zap.Error(err))
		}
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// VerifyServeEnvironment verifies the minimum required variables to run the serve command are present
func VerifyServeEnvironment() {
	// Verify all our required variables are set

	util.RequireVariables([]string{
		"BOT_TOKEN",
		"COMMAND_PREFIX",
		"OPENAI_TOKEN",
		"OPENAI_MODEL_IGUSTUS",
	})

}
