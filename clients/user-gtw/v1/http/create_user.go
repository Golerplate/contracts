package user_gtw_http_v1

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	user_gtw_v1_entities "github.com/golerplate/contracts/clients/user-gtw/v1/entities"
	"github.com/rs/zerolog/log"
)

func (c *UserGtwClient) CreateUser(ctx context.Context, user *user_gtw_v1_entities.CreateUserRequest) error {
	marshaledContent, err := json.Marshal(user)
	if err != nil {
		log.Error().Err(err).
			Msg("contracts.clients.user-gtw.v1.http.UserGtwClient.CreateUser: can not marshal request")
		return err
	}

	req, err := http.NewRequest("POST", c.baseURL+"internal/v1/users", strings.NewReader(string(marshaledContent)))
	if err != nil {
		log.Error().Err(err).
			Msg("contracts.clients.user-gtw.v1.http.UserGtwClient.CreateUser: can not create request")
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		log.Error().Err(err).
			Msg("contracts.clients.user-gtw.v1.http.UserGtwClient.CreateUser: can not do request")
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error().
			Str("status", resp.Status).
			Msg("contracts.clients.user-gtw.v1.http.UserGtwClient.CreateUser: unexpected status")
		return err
	}

	return nil
}
