CREATE TABLE "part" (
                        "id" serial NOT NULL PRIMARY KEY,
                        "manufacturer_id" integer NOT NULL,
                        "vendor_code" varchar(256) NOT NULL,
                        "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                        "deleted_at" TIMESTAMP WITH TIME ZONE
);

ALTER TABLE "part" ADD FOREIGN KEY ("manufacturer_id") REFERENCES "part_manufacturer" ("id");

INSERT INTO "part" ("manufacturer_id", "vendor_code")
VALUES (1, '1-3-3-7'), (1, '1-xxx-yyy'), (1, '1-zzz-ttt'), (2, '2-qwe-rty'), (3, '3-sss-aaa');