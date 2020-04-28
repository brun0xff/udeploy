package cfg

import (
	"log"
	"os"
	"strings"

	"github.com/turnerlabs/cstore/v4"
)

const (
	missingEnvOk = "MISSING_ENV_OK"

	url = "URL"
	env = "ENV"

	dbURI  = "DB_URI"
	dbName = "DB_NAME"

	oauthClientID           = "OAUTH_CLIENT_ID"
	oauthClientSecret       = "OAUTH_CLIENT_SECRET"
	oauthRedirectURL        = "OAUTH_REDIRECT_URL"
	oauthSignOutURL         = "OAUTH_SIGN_OUT_URL"
	oauthAuthURL            = "OAUTH_AUTH_URL"
	oauthTokenURL           = "OAUTH_TOKEN_URL"
	oauthSessSign           = "OAUTH_SESSION_SIGN"
	oauthScopes             = "OAUTH_SCOPES"
	oauthSignOutRedirectURL = "OAUTH_SIGN_OUT_REDIRECT_URL"

	sqsChangeQueue = "SQS_CHANGE_QUEUE"
	sqsAlarmQueue  = "SQS_ALARM_QUEUE"
	sqsS3Queue     = "SQS_S3_QUEUE"

	snsAlarmTopicArn = "SNS_ALARM_TOPIC_ARN"

	consoleLink = "CONSOLE_LINK"

	preCache = "PRE_CACHE"

	cstoreCatalog = "CSTORE_CATALOG"
	cstoreTags    = "CSTORE_TAGS"
)

// Get ...
var Get = map[string]string{
	preCache: "false",
}

func init() {

	fromEnv()

	if catalog, exists := os.LookupEnv(cstoreCatalog); exists {

		if err := fromRemote(catalog, strings.Split(os.Getenv(cstoreTags), ",")); err != nil {
			log.Fatal(err)
		}

		log.Printf("Loaded %s %s configuration.", catalog, os.Getenv(cstoreTags))
	}

	validate()
}

func validate() {
	errMsg := "environment variable %s required"

	_, missingEnvAllowed := os.LookupEnv(missingEnvOk)

	if _, exists := Get[url]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, url)
	}

	if _, exists := Get[env]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, env)
	}

	if _, exists := Get[dbURI]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, dbURI)
	}

	if _, exists := Get[dbName]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, dbName)
	}

	if _, exists := Get[oauthClientID]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, oauthClientID)
	}

	if _, exists := Get[oauthClientSecret]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, oauthClientSecret)
	}

	if _, exists := Get[oauthRedirectURL]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, oauthRedirectURL)
	}

	if _, exists := Get[oauthSignOutURL]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, oauthSignOutURL)
	}

	if _, exists := Get[oauthAuthURL]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, oauthAuthURL)
	}

	if _, exists := Get[oauthTokenURL]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, oauthTokenURL)
	}

	if _, exists := Get[oauthSessSign]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, oauthSessSign)
	}

	if _, exists := Get[sqsChangeQueue]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, sqsChangeQueue)
	}

	if _, exists := Get[sqsS3Queue]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, sqsS3Queue)
	}

	if _, exists := Get[sqsAlarmQueue]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, sqsAlarmQueue)
	}

	if _, exists := Get[snsAlarmTopicArn]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, snsAlarmTopicArn)
	}

	if _, exists := Get[consoleLink]; !missingEnvAllowed && !exists {
		log.Fatalf(errMsg, consoleLink)
	}
}

// Requires an "CSTORE_PATH" environment variable to be set
// supporting the token substitution in the config.yml file.
func fromRemote(catalog string, tags []string) error {
	config, err := cstore.PullEnv(catalog,
		cstore.Options{
			Tags:          tags,
			InjectSecrets: true,
		})

	if err != nil {
		return err
	}

	for k, v := range config {
		Get[k] = v
	}

	return nil
}

func fromEnv() {
	if v, exists := os.LookupEnv(url); exists {
		Get[url] = v
	}

	if v, exists := os.LookupEnv(env); exists {
		Get[env] = v
	}

	if v, exists := os.LookupEnv(dbURI); exists {
		Get[dbURI] = v
	}

	if v, exists := os.LookupEnv(dbName); exists {
		Get[dbName] = v
	}

	if v, exists := os.LookupEnv(oauthClientID); exists {
		Get[oauthClientID] = v
	}

	if v, exists := os.LookupEnv(oauthClientSecret); exists {
		Get[oauthClientSecret] = v
	}

	if v, exists := os.LookupEnv(oauthRedirectURL); exists {
		Get[oauthRedirectURL] = v
	}

	if v, exists := os.LookupEnv(oauthSignOutURL); exists {
		Get[oauthSignOutURL] = v
	}

	if v, exists := os.LookupEnv(oauthAuthURL); exists {
		Get[oauthAuthURL] = v
	}

	if v, exists := os.LookupEnv(oauthTokenURL); exists {
		Get[oauthTokenURL] = v
	}

	if v, exists := os.LookupEnv(oauthSessSign); exists {
		Get[oauthSessSign] = v
	}

	if v, exists := os.LookupEnv(oauthScopes); exists {
		Get[oauthScopes] = v
	}

	if v, exists := os.LookupEnv(oauthSignOutRedirectURL); exists {
		Get[oauthSignOutRedirectURL] = v
	}

	if v, exists := os.LookupEnv(sqsChangeQueue); exists {
		Get[sqsChangeQueue] = v
	}

	if v, exists := os.LookupEnv(sqsS3Queue); exists {
		Get[sqsS3Queue] = v
	}

	if v, exists := os.LookupEnv(sqsAlarmQueue); exists {
		Get[sqsAlarmQueue] = v
	}

	if v, exists := os.LookupEnv(snsAlarmTopicArn); exists {
		Get[snsAlarmTopicArn] = v
	}

	if v, exists := os.LookupEnv(consoleLink); exists {
		Get[consoleLink] = v
	}

	if v, exists := os.LookupEnv(preCache); exists {
		Get[preCache] = v
	}
}
