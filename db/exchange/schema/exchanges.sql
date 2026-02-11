BEGIN;

CREATE TYPE exchange_status AS ENUM ('pending', 'accepted', 'rejected', 'cancelled');

CREATE TABLE exchanges
(
    id         UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    from_ad_id varchar         NOT NULL,
    to_ad_id   varchar         NOT NULL,
    status     exchange_status NOT NULL DEFAULT 'pending',
    comment    TEXT,
    time       TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE INDEX idx_exchanges_status ON exchanges (status);

COMMIT;
