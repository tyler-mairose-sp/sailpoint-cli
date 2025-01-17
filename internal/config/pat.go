package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type PatConfig struct {
	ClientID     string    `mapstructure:"clientid"`
	ClientSecret string    `mapstructure:"clientsecret"`
	AccessToken  string    `mapstructure:"accesstoken"`
	Expiry       time.Time `mapstructure:"expiry"`
}

func GetPatToken() string {
	return viper.GetString("environments." + GetActiveEnvironment() + ".pat.accesstoken")
}

func SetPatToken(token string) {
	viper.Set("environments."+GetActiveEnvironment()+".pat.accesstoken", token)
}

func GetPatTokenExpiry() time.Time {
	return viper.GetTime("environments." + GetActiveEnvironment() + ".pat.expiry")
}

func SetPatTokenExpiry(expiry time.Time) {
	viper.Set("environments."+GetActiveEnvironment()+".pat.expiry", expiry)
}

func GetPatClientID() string {
	envSecret := os.Getenv("SAIL_CLIENT_ID")
	if envSecret != "" {
		return envSecret
	} else {
		return viper.GetString("environments." + GetActiveEnvironment() + ".pat.clientid")
	}
}

func GetPatClientSecret() string {
	envSecret := os.Getenv("SAIL_CLIENT_SECRET")
	if envSecret != "" {
		return envSecret
	} else {
		return viper.GetString("environments." + GetActiveEnvironment() + ".pat.clientsecret")
	}
}

func SetPatClientID(ClientID string) {
	viper.Set("environments."+GetActiveEnvironment()+".pat.clientid", ClientID)
}

func SetPatClientSecret(ClientSecret string) {
	viper.Set("environments."+GetActiveEnvironment()+".pat.clientsecret", ClientSecret)
}

func PATLogin() error {

	uri, err := url.Parse(GetTokenUrl())
	if err != nil {
		return err
	}

	query := &url.Values{}
	query.Add("grant_type", "client_credentials")
	uri.RawQuery = query.Encode()

	data := &url.Values{}
	data.Add("client_id", GetPatClientID())
	data.Add("client_secret", GetPatClientSecret())

	ctx := context.TODO()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to retrieve access token. status %s", resp.Status)
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var tResponse TokenResponse

	err = json.Unmarshal(raw, &tResponse)
	if err != nil {
		return err
	}

	now := time.Now()

	SetPatToken(tResponse.AccessToken)
	SetPatTokenExpiry(now.Add(time.Second * time.Duration(tResponse.ExpiresIn)))

	return nil
}
