package requests

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type NetworkByChainIdRequest struct {
	Id int64
}

func NewNetworkByChainIdRequest(r *http.Request) (*NetworkByChainIdRequest, error) {
	IdAsString, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get id from the url path")
	}

	return &NetworkByChainIdRequest{Id: int64(IdAsString)}, nil
}
