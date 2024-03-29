CREATE TABLE IF NOT EXISTS stock_location (
    id TEXT NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(6) NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,

    CONSTRAINT stock_location_pkey PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS stock_item (
    id TEXT NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(6) NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,

    CONSTRAINT stock_item_pkey PRIMARY KEY ("id")
);
