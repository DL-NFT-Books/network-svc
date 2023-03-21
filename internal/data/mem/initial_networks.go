package mem

import (
	"fmt"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"github.com/dl-nft-books/network-svc/internal/data"
	"reflect"
)

const yamlInitialNetworksKey = "initial-networks"

type InitialNetworker interface {
	InitialNetworks() []data.Network
}

type initialNetworker struct {
	getter kv.Getter
	once   comfig.Once

	initialNetworks []data.Network
}

func NewInitialNetworker(getter kv.Getter) InitialNetworker {
	return &initialNetworker{
		getter: getter,
	}
}

func (in *initialNetworker) InitialNetworks() []data.Network {
	in.readConfig()
	return in.initialNetworks
}

func (in *initialNetworker) readConfig() {
	in.once.Do(func() interface{} {
		cfg := struct {
			InitialNetworks []data.Network `fig:"data"`
		}{}
		err := figure.
			Out(&cfg).
			With(figure.BaseHooks, topHooks).
			From(kv.MustGetStringMap(in.getter, yamlInitialNetworksKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out signer"))
		}
		in.initialNetworks = cfg.InitialNetworks

		return nil
	})
}

var topHooks = figure.Hooks{
	"[]data.Network": func(value interface{}) (reflect.Value, error) {
		switch s := value.(type) {
		case []interface{}:
			networks := make([]data.Network, 0, len(s))
			var err error
			for _, rawElem := range s {
				mapElem, ok := rawElem.(map[interface{}]interface{})
				if !ok {
					return reflect.Value{}, errors.Wrap(err,
						"failed to cast mapElem to interface")
				}

				normMap := make(map[string]interface{}, len(mapElem))
				for key, value := range mapElem {
					strKey := fmt.Sprintf("%v", key)
					normMap[strKey] = value
				}

				var data data.Network

				err := figure.
					Out(&data).
					From(normMap).
					Please()
				if err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to figure out")
				}

				networks = append(networks, data)
			}

			return reflect.ValueOf(networks), nil
		default:
			return reflect.Value{}, errors.New("unexpected type while figure []data.Network")
		}
	},
}
