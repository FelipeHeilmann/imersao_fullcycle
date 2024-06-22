"use server"

import axios from "axios"
import { revalidateTag } from "next/cache"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"

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

export async function checkoutAction(prevState: any, {
    cardHash,
    email,
  }: {
    cardHash: string;
    email: string;
  }) {
    const cookieStore = cookies()
    const eventId = cookieStore.get("eventId")?.value
    const spots = JSON.parse(cookieStore.get("spots")?.value || "[]")
    const ticketKind = cookieStore.get("ticketKind")?.value || "full"
  
    try {
        await axios.post("http://localhost:8080/checkout", {
            eventId: eventId,
            cardHash: cardHash,
            ticketKind: ticketKind,
            spots,
            email,
        })

        revalidateTag(`events/${eventId}`)
        redirect(`/checkout/${eventId}/success`)

    }catch(e) {
        return { error: "Erro ao realizar a compra" }
    }   
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
  