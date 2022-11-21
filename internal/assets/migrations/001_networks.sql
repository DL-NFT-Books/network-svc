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
CREATE UNIQUE INDEX idx_networks_chain_id ON networks (chain_id);

-- +migrate Down
DROP INDEX IF EXISTS idx_networks_chain_id;
DROP TABLE networks;
