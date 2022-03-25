CREATE TABLE IF NOT EXISTS "integration" (
    "id" UUID PRIMARY KEY,
    "project_id" UUID NOT NULL,
    "client_platform_id" UUID REFERENCES "client_platform"("id"),
    "client_type_id" UUID REFERENCES "client_type"("id"),
    "role_id" UUID REFERENCES "role"("id"),
    "title" VARCHAR,
    "secret_key" VARCHAR(512),
    "ip_whitelist" JSONB,
    "active" SMALLINT,
    "expires_at" TIMESTAMP NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    UNIQUE ("project_id", "client_platform_id")
);

ALTER TABLE IF EXISTS "session" ADD COLUMN IF NOT EXISTS "integration_id" UUID REFERENCES "integration"("id");
ALTER TABLE IF EXISTS "session" ADD CONSTRAINT "integration_id_or_user_id" CHECK (
    ("user_id" IS NULL OR "integration_id" IS NULL) AND NOT
    ("user_id" IS NULL AND "integration_id" IS NULL)
);