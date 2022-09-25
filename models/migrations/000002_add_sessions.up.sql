CREATE TABLE sessions(
    session_id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    email varchar NOT NULL ,
    user_agent varchar NOT NULL,
    client_ip varchar NOT NULL,
    refresh_token varchar NOT NULL ,
    expires_at timestamptz NOT NULL ,
    is_blocked bool NOT NULL default false,
    created_at timestamptz NOT NULL DEFAULT NOW()
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("email") REFERENCES "accounts" ("owner_email");
