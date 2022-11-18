package requests

import (
	"encoding/json"
	"net/http"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/network-svc/resources"
)

var AddressRegexp = regexp.MustCompile("^(0x)?[0-9a-fA-F]{40}$")

type CreateNetworkRequest struct {
	Data resources.NetworkDetailed `json:"data"`
}

func NewCreateNetworkRequest(r *http.Request) (*CreateNetworkRequest, error) {
	var request CreateNetworkRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return &request, request.validate()
}

func (r CreateNetworkRequest) validate() error {
	return validation.Errors{
		"data/attributes/name": validation.Validate(r.Data.Attributes.Name, validation.Required),
		"data/attributes/chain_id": validation.Validate(
			r.Data.Attributes.ChainId,
			validation.Required,
			validation.Min(0)),
		"data/attributes/factory_address": validation.Validate(
			r.Data.Attributes.Name,
			validation.Required,
			validation.Match(AddressRegexp)),
		"data/attributes/rpc_url": validation.Validate(r.Data.Attributes.Name, validation.Required),
		"data/attributes/ws_url":  validation.Validate(r.Data.Attributes.Name, validation.Required),
	}.Filter()
}
