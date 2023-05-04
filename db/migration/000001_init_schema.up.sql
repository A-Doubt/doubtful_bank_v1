CREATE TABLE "account" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entry" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "account_id" uuid NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfer" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" uuid NOT NULL,
  "to_account_id" uuid NOT NULL,
  "amount" bigint NOT NULL,
  "user_id" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "note" varchar
);

CREATE INDEX ON "account" ("owner");

CREATE INDEX ON "transfer" ("from_account_id");

CREATE INDEX ON "transfer" ("to_account_id");

CREATE INDEX ON "transfer" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entry"."amount" IS 'can be positive or negative.';

COMMENT ON COLUMN "transfer"."amount" IS 'must be positive';

ALTER TABLE "entry" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("from_account_id") REFERENCES "account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("to_account_id") REFERENCES "account" ("id");
