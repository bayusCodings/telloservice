CREATE TABLE "verify_emails" (
  "id" bigserial PRIMARY KEY,
  "username" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "secret_code" VARCHAR(255) NOT NULL,
  "is_used" BOOLEAN NOT NULL DEFAULT false,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL,
  "expired_at" TIMESTAMPTZ NOT NULL DEFAULT (now() + interval '15 minutes')
);

ALTER TABLE "verify_emails" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "users" ADD COLUMN "is_email_verified" BOOLEAN NOT NULL DEFAULT false;
