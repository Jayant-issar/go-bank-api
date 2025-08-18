-- Rollback for 000002_fix: restore original column name
ALTER TABLE "entries" RENAME COLUMN "amount" TO "ammount";

-- Restore the original comment
COMMENT ON COLUMN "entries"."ammount" IS 'it can be negative or positive';
