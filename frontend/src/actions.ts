"use server"

import { cookies } from "next/headers"

export async function selectSpotAction(eventId: string, spotId: string) {
    const cookiesStore = cookies()
    
    const spots = JSON.parse(cookiesStore.get("spots")?.value || "[]")
    spots.push(spotId)
    const uniqueSpots = spots.filter((spot: string, index: number) => spots.indexOf(spot) === index)
    cookiesStore.set("spots", JSON.stringify(uniqueSpots))
    cookiesStore.set("eventId", eventId)
}

export async function unselectSpotAction(spotName: string) {
    const cookiesStore = cookies()

    const spots = JSON.parse(cookiesStore.get("spots")?.value || "[]")
    const newSpots = spots.filter((spot: string) => spot !== spotName )
    cookiesStore.set("spots", JSON.stringify(newSpots))
}

export async function selectTicketTypeAction(ticketKind: "full" | "half") {
    const cookieStore = cookies()
    cookieStore.set("ticketKind", ticketKind)
}

export async function clearSpotsAction() {
    const cookieStore = cookies()
    cookieStore.set("spots", "[]")
    cookieStore.set("eventId", "")
}
  