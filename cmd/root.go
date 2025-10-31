package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/fang"
	"github.com/charmbracelet/lipgloss"
	"github.com/marcelblijleven/gh-hookshot/internal/tui"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/keys"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/tuicontext"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	owner   string
	repo    string
	Version = "dev"
)

var rootCmd = &cobra.Command{
	Use:   "gh-hookshot",
	Short: "A TUI GitHub CLI extension for viewing repository webhook deliveries",
	Long: `A terminal ui extension for viewiing GitHub repository webhooks and
	deliveries.`,
	Run: func(cmd *cobra.Command, args []string) {
		lipgloss.SetHasDarkBackground(termenv.HasDarkBackground())

		ctx := &tuicontext.Context{
			Owner:   owner,
			Repo:    repo,
			Version: Version,
			Keys:    *keys.Keys,
		}
		model := tui.New(ctx)

		p := tea.NewProgram(
			model,
			tea.WithAltScreen(), // Opens full screen
		)

		if _, err := p.Run(); err != nil {
			log.Fatal("An error occurred while starting Hookshot", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := fang.Execute(context.Background(), rootCmd, fang.WithVersion(rootCmd.Version)); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Version = Version

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gh-hookshot.yaml)")
	rootCmd.PersistentFlags().StringVar(&owner, "owner", "", "specify the owner of the repository")
	rootCmd.PersistentFlags().StringVar(&repo, "repo", "", "specify the name of the repository")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".gh-hookshot" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gh-hookshot")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
