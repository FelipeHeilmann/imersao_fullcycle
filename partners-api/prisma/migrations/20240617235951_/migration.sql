/*
  Warnings:

  - You are about to drop the `Spot` table. If the table is not empty, all the data it contains will be lost.

*/
-- CreateEnum
CREATE TYPE "TicketKind" AS ENUM ('full', 'half');

-- DropForeignKey
ALTER TABLE "Spot" DROP CONSTRAINT "Spot_event_id_fkey";

-- DropTable
DROP TABLE "Spot";

-- CreateTable
CREATE TABLE "spots" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "event_id" TEXT NOT NULL,
    "status" "SpotStatus" NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "spots_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "tickets" (
    "id" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "ticket_kind" "TicketKind" NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "spotId" TEXT NOT NULL,

    CONSTRAINT "tickets_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "reservation_history" (
    "id" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "ticket_kind" "TicketKind" NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "spotId" TEXT NOT NULL,

    CONSTRAINT "reservation_history_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "tickets_spotId_key" ON "tickets"("spotId");

-- AddForeignKey
ALTER TABLE "spots" ADD CONSTRAINT "spots_event_id_fkey" FOREIGN KEY ("event_id") REFERENCES "events"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "tickets" ADD CONSTRAINT "tickets_spotId_fkey" FOREIGN KEY ("spotId") REFERENCES "spots"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "reservation_history" ADD CONSTRAINT "reservation_history_spotId_fkey" FOREIGN KEY ("spotId") REFERENCES "spots"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
