CREATE TABLE "accounts" (
    "id" BIGSERIAL PRIMARY KEY,
    "owner_email" varchar NOT NULL UNIQUE ,
    "balance" BIGINT NOT NULL,
    "currency" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "updated_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE "transactions" (
    "id" BIGSERIAL PRIMARY KEY,
    "transaction_id" BIGINT unique ,
    "from_account_id" BIGINT,
    "to_account_id" BIGINT,
    "amount" BIGINT NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE "account_transactions_entries" (
    "id" BIGSERIAL PRIMARY KEY,
    "transaction_id" BIGINT,
    "account_id" bigint,
    "other_account" bigint,
    "amount" bigint NOT NULL,
    "type" varchar,
    "created_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE INDEX ON "accounts" ("owner_email");

CREATE INDEX ON "transactions" ("from_account_id");

CREATE INDEX ON "transactions" ("to_account_id");

CREATE INDEX ON "transactions" ("from_account_id", "to_account_id");

CREATE INDEX ON "account_transactions_entries" ("account_id");

COMMENT ON COLUMN "transactions"."amount" IS 'can be negative depending on debit or credit';

COMMENT ON COLUMN "account_transactions_entries"."amount" IS 'must be positive';

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "account_transactions_entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

-- updated at timestamp function
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- setting trigger to update timestamp
CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON accounts
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();