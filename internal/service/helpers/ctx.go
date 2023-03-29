package helpers

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
	"github.com/dl-nft-books/doorman/connector"
	"github.com/dl-nft-books/network-svc/internal/data"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	networksQCtxKey
	doormanConnectorCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func CtxNetworksQ(entry data.NetworksQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, networksQCtxKey, entry)
	}
}

func NetworksQ(r *http.Request) data.NetworksQ {
	return r.Context().Value(networksQCtxKey).(data.NetworksQ).New()
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxDoormanConnector(entry connector.ConnectorI) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, doormanConnectorCtxKey, entry)
	}
}
func DoormanConnector(r *http.Request) connector.ConnectorI {
	return r.Context().Value(doormanConnectorCtxKey).(connector.ConnectorI)
}
