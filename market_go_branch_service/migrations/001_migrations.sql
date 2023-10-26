CREATE TABLE "branches" (
  "id" uuid PRIMARY KEY,
  "name" varchar,
  "address" varchar,
  "founded_at" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE TABLE "branch_products" (
  "product_id" uuid,
  "branch_id" uuid,
  "count" integer,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp
);

ALTER TABLE "branch_products" ADD FOREIGN KEY ("branch_id") REFERENCES "branches" ("id");