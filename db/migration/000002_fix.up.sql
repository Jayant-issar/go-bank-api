-- Fix typo in entries table: rename column "ammount" to "amount"
ALTER TABLE "entries" RENAME COLUMN "ammount" TO "amount";

-- Ensure the column comment remains correct after rename
COMMENT ON COLUMN "entries"."amount" IS 'it can be negative or positive';
