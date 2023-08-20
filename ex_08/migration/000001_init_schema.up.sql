CREATE TABLE "users"
(
    "username"            varchar PRIMARY KEY,
    "full_name"           varchar        NOT NULL,
    "email"               varchar UNIQUE NOT NULL,
    "hashed_password"     varchar        NOT NULL,
    "salt"                varchar        NOT NULL,
    "password_changed_at" timestamptz    NOT NULL DEFAULT '0001-01-01',
    "created_at"          timestamptz    NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions"
(
    "id"            bigserial PRIMARY KEY,
    "username"      varchar     NOT NULL,
    "refresh_token" varchar     NOT NULL,
    "user_agent"    varchar     NOT NULL,
    "client_ip"     varchar     NOT NULL,
    "is_blocked"    boolean     NOT NULL DEFAULT false,
    "expires_at"    timestamptz NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "brands"
(
    "id"               bigserial PRIMARY KEY,
    "code"             varchar UNIQUE NOT NULL,
    "name"             varchar UNIQUE NOT NULL,
    "website"          varchar,
    "active"           boolean,
    "description"      text,
    "creator_username" varchar        NOT NULL,
    "created_at"       timestamptz    NOT NULL DEFAULT (now()),
    "updated_at"       timestamptz    NOT NULL DEFAULT (now())
);

CREATE TABLE "categories"
(
    "id"               bigserial PRIMARY KEY,
    "parent_id"        bigserial,
    "name"             varchar UNIQUE NOT NULL,
    "description"      text,
    "creator_username" varchar        NOT NULL,
    "created_at"       timestamptz    NOT NULL DEFAULT (now()),
    "updated_at"       timestamptz    NOT NULL DEFAULT (now())
);

CREATE TABLE "products"
(
    "id"               bigserial PRIMARY KEY,
    "brand_id"         bigserial,
    "category_id"      bigserial,
    "code"             varchar UNIQUE NOT NULL,
    "name"             varchar UNIQUE NOT NULL,
    "quantity"         int,
    "weight"           bigint         NOT NULL,
    "weight_unit"      varchar        NOT NULL,
    "price"            bigint,
    "currency"         varchar        NOT NULL,
    "description"      text,
    "slug"             varchar UNIQUE,
    "sold"             int,
    "rate"             int,
    "reviews"          int,
    "status"           varchar        NOT NULL,
    "old_status"       varchar        NOT NULL,
    "creator_username" varchar        NOT NULL,
    "created_at"       timestamptz    NOT NULL DEFAULT (now()),
    "updated_at"       timestamptz    NOT NULL DEFAULT (now()),
    "deleted_at"       timestamptz
);

CREATE TABLE "carts"
(
    "id"             bigserial PRIMARY KEY,
    "code"           varchar UNIQUE NOT NULL,
    "weight"         bigint         NOT NULL,
    "weight_unit"    varchar        NOT NULL,
    "quantity"       int,
    "total"          bigint,
    "currency"       varchar        NOT NULL,
    "note"           text,
    "active"         boolean,
    "owner_username" varchar        NOT NULL,
    "created_at"     timestamptz    NOT NULL DEFAULT (now()),
    "updated_at"     timestamptz    NOT NULL DEFAULT (now())
);

CREATE TABLE "cart_items"
(
    "cart_id"    bigserial   NOT NULL,
    "product_id" bigserial   NOT NULL,
    "quantity"   int,
    "price"      bigint,
    "total"      bigint,
    "currency"   varchar     NOT NULL,
    "note"       text,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "payments"
(
    "id"             bigserial PRIMARY KEY,
    "cart_id"        bigserial   NOT NULL,
    "discount"       bigint,
    "total"          bigint,
    "currency"       varchar     NOT NULL,
    "status"         varchar     NOT NULL,
    "note"           text,
    "owner_username" varchar     NOT NULL,
    "created_at"     timestamptz NOT NULL DEFAULT (now()),
    "updated_at"     timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "brands" ("code");

CREATE INDEX ON "brands" ("name");

CREATE INDEX ON "categories" ("name");

CREATE INDEX ON "categories" ("parent_id");

CREATE INDEX ON "products" ("code");

CREATE INDEX ON "products" ("name");

CREATE INDEX ON "carts" ("code");

CREATE INDEX ON "cart_items" ("cart_id");

CREATE UNIQUE INDEX ON "cart_items" ("cart_id", "product_id");

CREATE INDEX ON "payments" ("cart_id");

ALTER TABLE "sessions"
    ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "brands"
    ADD FOREIGN KEY ("creator_username") REFERENCES "users" ("username");

ALTER TABLE "categories"
    ADD FOREIGN KEY ("parent_id") REFERENCES "categories" ("id");

ALTER TABLE "categories"
    ADD FOREIGN KEY ("creator_username") REFERENCES "users" ("username");

ALTER TABLE "products"
    ADD FOREIGN KEY ("creator_username") REFERENCES "users" ("username");

ALTER TABLE "carts"
    ADD FOREIGN KEY ("owner_username") REFERENCES "users" ("username");

ALTER TABLE "cart_items"
    ADD FOREIGN KEY ("cart_id") REFERENCES "carts" ("id");

ALTER TABLE "cart_items"
    ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "payments"
    ADD FOREIGN KEY ("cart_id") REFERENCES "carts" ("id");

ALTER TABLE "payments"
    ADD FOREIGN KEY ("owner_username") REFERENCES "users" ("username");
