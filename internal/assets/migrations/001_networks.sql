-- +migrate Up
CREATE TABLE networks
(
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL DEFAULT '',
    chain_id BIGINT NOT NULL DEFAULT 0,
    rpc_url text NOT NULL DEFAULT '',
    ws_url text NOT NULL DEFAULT '',
    factory_address VARCHAR(42) NOT NULL DEFAULT ''
);

-- +migrate Down
DROP TABLE networks;
