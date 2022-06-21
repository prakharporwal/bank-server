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
    "transaction_id" BIGINT unique NOT NULL ,
    "from_account_id" BIGINT NOT NULL ,
    "to_account_id" BIGINT not null ,
    "amount" BIGINT NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE "account_transactions_entries" (
    "id" BIGSERIAL PRIMARY KEY,
    "transaction_id" BIGINT NOT NULL ,
    "account_id" bigint NOT NULL ,
    "other_account" bigint NOT NULL ,
    "amount" bigint NOT NULL,
    "type" varchar NOT NULL ,
    "created_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE "users"(
    user_id BIGSERIAL PRIMARY KEY ,
    username varchar NOT NULL ,
    user_email varchar UNIQUE NOT NULL ,
    password_hash varchar NOT NULL ,
    is_verified boolean NOT NULL DEFAULT false,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX ON "accounts" ("owner_email");

CREATE INDEX ON "transactions" ("from_account_id");

CREATE INDEX ON "transactions" ("to_account_id");

CREATE INDEX ON "transactions" ("from_account_id", "to_account_id");

CREATE INDEX ON "account_transactions_entries" ("account_id");

CREATE INDEX ON "users" ("user_email");

CREATE INDEX ON "users" ("username");

COMMENT ON COLUMN "transactions"."amount" IS 'can be negative depending on debit or credit';

COMMENT ON COLUMN "account_transactions_entries"."amount" IS 'must be positive';

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "account_transactions_entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

-- ALTER TABLE "users"  FOREIGN KEY ("user_email") REFERENCES "accounts" ("owner_email");


-- updated at timestamp function
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- setting trigger to update timestamp accounts table
CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON accounts
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


-- setting trigger to update timestamp user table
CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON users
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();