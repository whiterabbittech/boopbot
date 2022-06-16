package contribution

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/whiterabbittech/boopbot/config"
	"github.com/whiterabbittech/boopbot/givebutter"
)

var conf *config.Config

func init() {
	conf = config.NewFromEnv()
	setupLogger(conf)
	givebutter.SetAPIKey(conf.GivebutterAPIKey())
}

type ScanContributionsReq struct{}

type ScanContributionsResp struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(req ScanContributionsReq) (*ScanContributionsResp, error) {
	logrus.Info("Received Request to Scan GiveButter Contributions")
	var ctx = context.Background()
	var txns, err = givebutter.GetTransactions(ctx, givebutter.All)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	logrus.WithFields(logrus.Fields{
		"first_name": txns.Data[0].FirstName,
		"last_name":  txns.Data[0].LastName,
		"donated":    txns.Data[0].Donated,
	}).Info("API Call Sent to GiveButter")
	var resp = &ScanContributionsResp{
		StatusCode: 200,
		Body:       "Foobar",
	}

	return resp, nil
}

// setupLogger will configure Logrus logging using
// configuration from the environment.
func setupLogger(conf *config.Config) {
	// To set up the logger, we decide whether we want
	// JSON logging or human-oriented logging.
	// We also set the log level.
	var formatter logrus.Formatter
	// Only use text-based logging when developing locally.
	// If this app is deployed to any other environment,
	// use JSON logging.
	if conf.Env() == config.None {
		formatter = &logrus.TextFormatter{
			ForceColors: true,
		}
	} else {
		formatter = &logrus.JSONFormatter{}
	}
	logrus.SetFormatter(formatter)
	logrus.SetLevel(conf.LogLevel())
}
