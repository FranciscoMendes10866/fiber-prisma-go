# Migration `20200901161041-init`

This migration has been generated at 9/1/2020, 4:10:41 PM.
You can check out the [state of the schema](./schema.prisma) after the migration.

## Database Steps

```sql
CREATE TABLE "public"."User" (
"age" integer   ,"createdAt" timestamp(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP,"email" text  NOT NULL ,"id" text  NOT NULL ,"name" text   ,
    PRIMARY KEY ("id"))

CREATE TABLE "public"."Post" (
"authorID" text  NOT NULL ,"content" text   ,"createdAt" timestamp(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP,"id" text  NOT NULL ,"published" boolean  NOT NULL ,"title" text  NOT NULL ,"updatedAt" timestamp(3)  NOT NULL ,
    PRIMARY KEY ("id"))

CREATE UNIQUE INDEX "User.email" ON "public"."User"("email")

ALTER TABLE "public"."Post" ADD FOREIGN KEY ("authorID")REFERENCES "public"."User"("id") ON DELETE CASCADE  ON UPDATE CASCADE
```

## Changes

```diff
diff --git schema.prisma schema.prisma
migration 20200901153503-init..20200901161041-init
--- datamodel.dml
+++ datamodel.dml
@@ -1,7 +1,7 @@
 datasource db {
-    provider = "sqlite"
-    url = "***"
+  provider = "postgresql"
+  url      = "postgresql://golang:golang@localhost:9000/golang"
 }
 generator db {
     provider = "go run github.com/prisma/prisma-client-go"
```


