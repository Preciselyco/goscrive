package scrive_test

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"testing"

	scrive "github.com/Preciselyco/goscrive"
	"github.com/joho/godotenv"
)

var envOnce sync.Once

func readEnv() {
	envOnce.Do(func() {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	})
}

func getConfig() (scrive.Config, error) {
	readEnv()
	config := scrive.Config{
		PAC: &scrive.PAC{
			ClientCredentialsIdentifier: os.Getenv("CLIENT_CREDENTIALS_IDENTIFIER"),
			ClientCredentialsSecret:     os.Getenv("CLIENT_CREDENTIALS_SECRET"),
			TokenCredentialsIdentifier:  os.Getenv("TOKEN_CREDENTIALS_IDENTIFIER"),
			TokenCredentialsSecret:      os.Getenv("TOKEN_CREDENTIALS_SECRET"),
		},
	}
	return config, nil
}

func getClient() *scrive.Client {
	config, err := getConfig()
	if err != nil {
		panic(err)
	}
	cli, err := scrive.NewClient(config)
	if err != nil {
		panic(err)
	}
	return cli
}

func marshalIndentFail(t *testing.T, context string, v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	failIfE(t, context, err)
	return string(b)
}

func dumps(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return ""
	}
	return string(b)
}

func failLog(t *testing.T, context string, v interface{}) {
	vStr := dumps(v)
	fmt.Printf("%s: \n%s\n", context, vStr)
	t.Fail()
}

func failIfE(t *testing.T, context string, err error) {
	if err != nil {
		failLog(t, context, err)
	}
}

func failIfScriveE(t *testing.T, context string, err *scrive.ScriveError) {
	if err != nil {
		failLog(t, context, err)
	}
}
