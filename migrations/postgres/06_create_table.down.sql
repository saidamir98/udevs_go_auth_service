ALTER TABLE IF EXISTS "session" DROP CONSTRAINT "only_one_value";
ALTER TABLE IF EXISTS "session" DROP COLUMN IF EXISTS "integration_id";

DROP TABLE IF EXISTS "integration";
