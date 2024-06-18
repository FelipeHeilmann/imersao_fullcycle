/*
  Warnings:

  - Added the required column `status` to the `reservation_history` table without a default value. This is not possible if the table is not empty.

*/
-- CreateEnum
CREATE TYPE "TicketStatus" AS ENUM ('reserved', 'canceled');

-- AlterTable
ALTER TABLE "reservation_history" ADD COLUMN     "status" "TicketStatus" NOT NULL;
