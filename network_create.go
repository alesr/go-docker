package docker // import "docker.io/go-docker"

import (
	"encoding/json"

	"docker.io/go-docker/api/types"
	"golang.org/x/net/context"
)

// NetworkCreate creates a new network in the docker host.
func (cli *Client) NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error) {
	networkCreateRequest := types.NetworkCreateRequest{
		NetworkCreate: options,
		Name:          name,
	}
	var response types.NetworkCreateResponse
	serverResp, err := cli.post(ctx, "/networks/create", nil, networkCreateRequest, nil)
	if err != nil {
		return response, err
	}

	json.NewDecoder(serverResp.body).Decode(&response)
	ensureReaderClosed(serverResp)
	return response, err
}
