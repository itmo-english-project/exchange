BEGIN;

CREATE TYPE exchange_status AS ENUM ('pending', 'accepted', 'rejected', 'cancelled');

CREATE TABLE exchanges
(
    id         UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    from_ad_id UUID            NOT NULL,
    to_ad_id   UUID            NOT NULL,
    status     exchange_status NOT NULL DEFAULT 'pending',
    comment    TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE INDEX idx_exchanges_status ON exchanges (status);

COMMIT;
