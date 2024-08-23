CREATE TABLE "accounts" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" BIGINT NOT NULL,
  "balance" DECIMAL(12, 2) NOT NULL,
  "currency" VARCHAR(50) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "transactions" (
  "id" BIGSERIAL PRIMARY KEY,
  "transaction_id" VARCHAR(255) NOT NULL,
  "from_account_id" BIGINT NOT NULL,
  "to_account_id" BIGINT NOT NULL,
  "amount" DECIMAL(12, 2) NOT NULL,
  "status" VARCHAR(100) NOT NULL,
  "description" VARCHAR(255) NULL,
  "transaction_type" VARCHAR(100) NOT NULL,
  "transaction_reference" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE INDEX ON "accounts" ("user_id");

CREATE INDEX ON "transactions" ("from_account_id");

CREATE INDEX ON "transactions" ("to_account_id");

CREATE INDEX ON "transactions" ("from_account_id", "to_account_id");

CREATE INDEX ON "transactions" ("transaction_id");

CREATE INDEX ON "transactions" ("transaction_reference");

COMMENT ON COLUMN "transactions"."amount" IS 'must be positive';

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
