import { EventModel } from "@/models"
import Link from "next/link"
import { EventImage } from "./EventImage"

export type CardProps = {
    event: EventModel
}

export default function EventCard({ event }: CardProps) {
    return (
        <Link href={`/event/${event.id}/spots`}>
            <div className="flex w-[277px] flex-col rounded-2xl bg-secondary">
                <EventImage src={event.imageUrl} alt={event.name}/>
                <div className="flex flex-col gap-y-2 px-4 py-6">
                    <p className="text-sm uppercase text-subtitle">
                        {new Date(event.date).toLocaleDateString("pt-Br", {
                            weekday: "long",
                            day: "2-digit",
                            month: "2-digit",
                            year: "numeric"
                        })}
                    </p>
                    <p className="font-semibold">{event.name}</p>
                    <p className="text-sm font-normal">{event.location}</p>
                </div>
            </div>
        </Link>
    )
}