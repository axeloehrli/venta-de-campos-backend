CREATE TABLE "usuarios" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "nombre_usuario" varchar NOT NULL,
  "nombre" varchar NOT NULL,
  "apellido" varchar NOT NULL,
  "email" varchar NOT NULL,
  "fecha_creacion" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "campos" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "id_usuario" bigint NOT NULL,
  "titulo" varchar NOT NULL,
  "tipo" varchar NOT NULL,
  "hectareas" bigint NOT NULL,
  "precio_por_hectaria" float NOT NULL,
  "ciudad" varchar NOT NULL,
  "provincia" varchar NOT NULL
);

CREATE INDEX ON "usuarios" ("nombre_usuario");

CREATE INDEX ON "campos" ("id");

CREATE INDEX ON "campos" ("provincia");

ALTER TABLE "campos" ADD FOREIGN KEY ("id_usuario") REFERENCES "usuarios" ("id");
