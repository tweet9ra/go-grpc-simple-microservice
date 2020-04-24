CREATE TABLE "part_manufacturer" (
  "id" serial NOT NULL PRIMARY KEY,
  "name" varchar(256) NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO "part_manufacturer" ("name", "created_at")
VALUES ('Производитель 1', 'now()'), ('Производитель 2', 'now()'), ('Производитель 3', 'now()');