import EventCard from "@/components/EventCard";
import Title from "@/components/Title";
import { EventModel } from "@/models";
import axios from "axios";

export async function getEvents(): Promise<EventModel[]> {
  const response = await axios.get("http://localhost:8080/events")

  return response.data.events
}

export default async function Home() {
  const events = await getEvents()
  return (
    <main className="mt-10 flex flex-col">
      <Title>Eventos disponíveis</Title>
      <div className="mt-8 sm:grid sm:grid-cols-auto-fit-cards flex flex-wrap justify-center gap-x-2 gap-y-4">
        {events.map((event) => <EventCard key={event.id} event={event}></EventCard>)}
      </div>
    </main>
  )
}
