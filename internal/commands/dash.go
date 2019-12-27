package commands

import (
	"context"
	"flag"
	"fmt"
	"github.com/kubenext/kubefun/internal/log"
	"github.com/kubenext/kubefun/pkg/dash"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	golog "log"
	"os"
	"os/signal"
	"strings"
)

func newKubefunCmd(version string) *cobra.Command {
	var verboseLevel int

	kubefunCmd := &cobra.Command{
		Use:   "kubefun",
		Short: "kubefun kubernetes dashboard",
		Long:  "kubefun is a dashboard for high bandwidth cluster analysis operations",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// Count flags don't work correctly when accessed directly from viper so we must use CountVarP and set the result.
			viper.Set("verbosity", verboseLevel)

			z, err := newZapLogger(viper.GetInt("verbosity"))
			if err != nil {
				golog.Printf("failed to initialize logger: %v", err)
				os.Exit(1)
			}
			defer func() {
				// this fails, but it should be safe to ignore according
				// to https://github.com/uber-go/zap/issues/328
				_ = z.Sync()
			}()

			logger := log.Wrap(z.Sugar())

			sigCh := make(chan os.Signal, 1)
			signal.Notify(sigCh, os.Interrupt)

			runCh := make(chan bool, 1)

			shutdownCh := make(chan bool, 1)

			err = bindViper(cmd)
			if err != nil {
				logger.WithErr(err).Errorf("unable to bind flags")
				os.Exit(1)
			}

			logger.Debugf("disable-open-browser: %s", viper.Get("disable-open-browser"))

			if viper.GetString("kubeconfig") == "" {
				viper.Set("kubeconfig", clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename())
			}

			go func() {
				options := dash.Options{
					DisableClusterOverview: viper.GetBool("disable-cluster-overview"),
					EnableOpenCensus:       viper.GetBool("enable-opencensus"),
					KubeConfig:             viper.GetString("kubeconfig"),
					Namespace:              viper.GetString("namespace"),
					FrontendURL:            viper.GetString("ui-url"),
					Context:                viper.GetString("context"),
					ClientQPS:              float32(viper.GetFloat64("client-qps")),
					ClientBurst:            viper.GetInt("client-burst"),
					UserAgent:              fmt.Sprintf("kubefun/%s", version),
				}

				klogVerbosity := viper.GetInt("klog-verbosity")
				if klogVerbosity > 0 {
					klog.InitFlags(nil)
					verbosityOpt := fmt.Sprintf("-v=%d", klogVerbosity)
					if err := flag.CommandLine.Parse([]string{verbosityOpt, "-logtostderr=true"}); err != nil {
						logger.WithErr(err).Errorf("unable to parse klog flags")
					}

				}

				if err := dash.Run(ctx, logger, shutdownCh, options); err != nil {
					logger.WithErr(err).Errorf("dashboard failed")
					os.Exit(1)
				}

				runCh <- true
			}()

			select {
			case <-sigCh:
				logger.Debugf("Shutting dashboard down due to interrupt")
				cancel()
				// TODO implement graceful shutdown semantics (GH#494)

				<-shutdownCh
			case <-runCh:
				logger.Debugf("Dashboard has exited")
			}
		},
	}

	// All flags can also be environment variables by adding the KUBEFUN_ prefix
	// and replacing - with _. Example: KUBEFUN_DISABLE_CLUSTER_OVERVIEW
	kubefunCmd.Flags().SortFlags = false

	kubefunCmd.Flags().StringP("context", "", "", "initial context")
	kubefunCmd.Flags().BoolP("disable-cluster-overview", "", false, "disable cluster overview")
	kubefunCmd.Flags().BoolP("enable-feature-applications", "", false, "enable applications feature")
	kubefunCmd.Flags().String("kubeconfig", "", "absolute path to kubeConfig file")
	kubefunCmd.Flags().StringP("namespace", "n", "", "initial namespace")
	kubefunCmd.Flags().StringP("plugin-path", "", "", "plugin path")
	kubefunCmd.Flags().CountVarP(&verboseLevel, "verbosity", "v", "verbosity level")

	kubefunCmd.Flags().StringP("accepted-hosts", "", "", "accepted hosts list [DEV]")
	kubefunCmd.Flags().Float32P("client-qps", "", 200, "maximum QPS for client [DEV]")
	kubefunCmd.Flags().IntP("client-burst", "", 400, "maximum burst for client throttle [DEV]")
	kubefunCmd.Flags().BoolP("disable-open-browser", "", false, "disable automatic launching of the browser [DEV]")
	kubefunCmd.Flags().BoolP("enable-opencensus", "c", false, "enable open census [DEV]")
	kubefunCmd.Flags().IntP("klog-verbosity", "", 0, "klog verbosity level [DEV]")
	kubefunCmd.Flags().StringP("listener-addr", "", "", "listener address for the kubefun frontend [DEV]")
	kubefunCmd.Flags().StringP("local-content", "", "", "local content path [DEV]")
	kubefunCmd.Flags().StringP("proxy-frontend", "", "", "url to send frontend request to [DEV]")
	kubefunCmd.Flags().String("ui-url", "", "dashboard url [DEV]")

	return kubefunCmd
}

func bindViper(cmd *cobra.Command) error {
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("KUBEFUN")
	viper.AutomaticEnv()

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}
	if err := viper.BindEnv("xdg-config-home", "XDG_CONFIG_HOME"); err != nil {
		return err
	}
	if err := viper.BindEnv("home", "HOME"); err != nil {
		return err
	}
	if err := viper.BindEnv("local-app-data", "LOCALAPPDATA"); err != nil {
		return err
	}
	if err := viper.BindEnv("kubeconfig", "KUBECONFIG"); err != nil {
		return err
	}

	return nil
}

// Returns a new zap logger, setting level according to the provided
// verbosity level as an offset of the base level, Info.
// i.e. verboseLevel==0, level==Info
//      verboseLevel==1, level==Debug
func newZapLogger(verboseLevel int) (*zap.Logger, error) {
	level := zapcore.InfoLevel - zapcore.Level(verboseLevel)
	if level < zapcore.DebugLevel || level > zapcore.FatalLevel {
		level = zapcore.DebugLevel
	}

	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	return cfg.Build()
}
