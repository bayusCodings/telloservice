CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" VARCHAR(255) NOT NULL UNIQUE,
  "password" VARCHAR(255) NOT NULL,
  "first_name" VARCHAR(150) NOT NULL,
  "last_name" VARCHAR(150) NOT NULL,
  "email" VARCHAR(255) NOT NULL UNIQUE,
  "password_changed_at" TIMESTAMPTZ,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

-- CREATE UNIQUE INDEX ON "accounts" ("user_id", "currency");
ALTER TABLE "accounts" ADD CONSTRAINT "user_id_currency_key" UNIQUE ("user_id", "currency");
