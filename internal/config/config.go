package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/square/go-jose.v2/jwt"

	"github.com/fatih/color"
	sailpoint "github.com/sailpoint-oss/golang-sdk"
	"github.com/sailpoint-oss/sailpoint-cli/internal/log"
	"github.com/sailpoint-oss/sailpoint-cli/internal/types"
	"github.com/spf13/viper"
)

var ErrAccessTokenExpired = fmt.Errorf("accesstoken is expired")

const (
	configFolder   = ".sailpoint"
	configYamlFile = "config.yaml"
)

type Token struct {
	AccessToken string    `mapstructure:"accesstoken"`
	Expiry      time.Time `mapstructure:"expiry"`
}

type Environment struct {
	TenantURL string    `mapstructure:"tenanturl"`
	BaseURL   string    `mapstructure:"baseurl"`
	Pat       PatConfig `mapstructure:"pat"`
	OAuth     Token     `mapstructure:"oauth"`
}

type CLIConfig struct {

	//Standard Variables
	ExportTemplatesPath string                 `mapstructure:"exporttemplatespath"`
	SearchTemplatesPath string                 `mapstructure:"searchtemplatespath"`
	Debug               bool                   `mapstructure:"debug"`
	AuthType            string                 `mapstructure:"authtype"`
	ActiveEnvironment   string                 `mapstructure:"activeenvironment"`
	Environments        map[string]Environment `mapstructure:"environments"`

	//Pipline Variables
	ClientID     string    `mapstructure:"clientid, omitempty"`
	ClientSecret string    `mapstructure:"clientsecret, omitempty"`
	BaseURL      string    `mapstructure:"base_url, omitempty"`
	AccessToken  string    `mapstructure:"accesstoken"`
	Expiry       time.Time `mapstructure:"expiry"`
}

func GetCustomSearchTemplatePath() string {
	return viper.GetString("searchtemplatespath")
}

func GetCustomExportTemplatePath() string {
	return viper.GetString("exporttemplatespath")
}

func SetCustomSearchTemplatePath(customsearchtemplatespath string) {
	viper.Set("searchtemplatespath", customsearchtemplatespath)
}

func SetCustomExportTemplatePath(customsearchtemplatespath string) {
	viper.Set("exporttemplatespath", customsearchtemplatespath)
}

func GetEnvironments() map[string]interface{} {
	return viper.GetStringMap("environments")
}

func GetAuthType() string {
	return strings.ToLower(viper.GetString("authtype"))
}

func SetAuthType(AuthType string) {
	viper.Set("authtype", strings.ToLower(AuthType))
}

func GetDebug() bool {
	return viper.GetBool("debug")
}

func SetDebug(Debug bool) {
	viper.Set("debug", Debug)
}

func GetActiveEnvironment() string {
	return strings.ToLower(viper.GetString("activeenvironment"))
}

func SetActiveEnvironment(activeEnv string) {
	viper.Set("activeenvironment", strings.ToLower(activeEnv))
}

func InitConfig() error {

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	viper.AddConfigPath(filepath.Join(home, ".sailpoint"))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("sail")

	viper.SetDefault("authtype", "pat")
	viper.SetDefault("exporttemplatespath", "")
	viper.SetDefault("searchtemplatespath", "")
	viper.SetDefault("debug", false)
	viper.SetDefault("activeenvironment", "default")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			// IGNORE they may be using env vars
		} else {
			// Config file was found but another error was produced
			return err
		}
	}

	return nil
}

func InitAPIClient() (*sailpoint.APIClient, error) {
	var apiClient *sailpoint.APIClient

	err := Validate()
	if err != nil {
		return apiClient, err
	}

	token, err := GetAuthToken()
	if err != nil && GetDebug() {
		color.Yellow("unable to retrieve accesstoken: %s ", err)
	}

	configuration := sailpoint.NewConfiguration(sailpoint.ClientConfiguration{Token: token, BaseURL: GetBaseUrl()})
	apiClient = sailpoint.NewAPIClient(configuration)
	if !GetDebug() {
		var DevNull types.DevNull
		apiClient.V3.GetConfig().HTTPClient.Logger = DevNull
		apiClient.Beta.GetConfig().HTTPClient.Logger = DevNull
	}

	return apiClient, nil
}

func CheckToken(tokenString string) error {
	var claims map[string]interface{}

	token, err := jwt.ParseSigned(tokenString)
	if err != nil {
		return err
	}

	token.UnsafeClaimsWithoutVerification(&claims)

	if claims["user_name"] == nil {
		log.Log.Warn("It looks like the token you are using is missing a user context, this will cause many of the CLI commands to fail.")
	}

	if GetDebug() {
		log.Log.Info("Token Debug Info", "user_name", claims["user_name"], "org", claims["org"], "pod", claims["pod"])
	}

	return nil
}

func GetAuthToken() (string, error) {

	var token string

	err := InitConfig()
	if err != nil {
		return "", err
	}

	err = Validate()
	if err != nil {
		return "", err
	}

	switch GetAuthType() {
	case "pat":
		if GetPatTokenExpiry().After(time.Now()) {

			token = GetPatToken()

		} else {

			err = PATLogin()
			if err != nil {
				return "", err
			}

			token = GetPatToken()
		}
	case "oauth":
		return "", fmt.Errorf("oauth is not currently supported")
		// if GetOAuthTokenExpiry().After(time.Now()) {
		// 	return GetOAuthToken(), nil
		// } else {
		// 	err = OAuthLogin()
		// 	if err != nil {
		// 		return "", err
		// 	}

		// 	return GetOAuthToken(), nil
		// }
	default:
		return "", fmt.Errorf("invalid authtype configured")
	}

	err = CheckToken(token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetBaseUrl() string {
	envBaseUrl := os.Getenv("SAIL_BASE_URL")
	if envBaseUrl != "" {
		return envBaseUrl
	} else {
		return viper.GetString("environments." + GetActiveEnvironment() + ".baseurl")
	}
}

func GetTenantUrl() string {
	return viper.GetString("environments." + GetActiveEnvironment() + ".tenanturl")
}

func SetBaseUrl(baseUrl string) {
	viper.Set("environments."+GetActiveEnvironment()+".baseurl", baseUrl)
}

func SetTenantUrl(tenantUrl string) {
	viper.Set("environments."+GetActiveEnvironment()+".tenanturl", tenantUrl)
}

func GetTokenUrl() string {
	return GetBaseUrl() + "/oauth/token"
}

func GetAuthorizeUrl() string {
	return GetTenantUrl() + "/oauth/authorize"
}

func GetConfig() (CLIConfig, error) {
	var Config CLIConfig

	err := viper.Unmarshal(&Config)
	if err != nil {
		return Config, err
	}
	return Config, nil
}

func SaveConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	if _, err := os.Stat(filepath.Join(home, configFolder)); os.IsNotExist(err) {
		err = os.Mkdir(filepath.Join(home, configFolder), 0777)
		if err != nil {
			log.Log.Warn("failed to create %s folder for config. %v", configFolder, err)
		}
	}

	err = viper.WriteConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = viper.SafeWriteConfig()
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func Validate() error {
	var errors int
	authType := GetAuthType()

	switch authType {

	case "pat":

		if GetBaseUrl() == "" {
			log.Log.Error("configured environment is missing BaseURL")
			errors++
		}

		if GetPatClientID() == "" {
			log.Log.Error("configured environment is missing PAT ClientID")
			errors++
		}

		if GetPatClientSecret() == "" {
			log.Log.Error("configured environment is missing PAT ClientSecret")
			errors++
		}

	case "oauth":

		log.Log.Error("oauth is not currently supported")
		errors++

		// if config.Environments[config.ActiveEnvironment].BaseURL == "" {
		// 	return fmt.Errorf("configured environment is missing BaseURL")
		// }

		// if config.Environments[config.ActiveEnvironment].TenantURL == "" {
		// 	return fmt.Errorf("configured environment is missing TenantURL")
		// }

	default:

		log.Log.Error("invalid authtype '%s' configured", authType)
		errors++

	}

	if errors > 0 {
		return fmt.Errorf("configuration invalid, errors: %v", errors)
	}

	return nil
}
