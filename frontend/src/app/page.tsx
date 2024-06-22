import EventCard from "@/components/EventCard";
import Title from "@/components/Title";
import { EventModel } from "@/models";
import Image from "next/image";

export default function Home() {
  const events: EventModel[] = [
    {
      id: "1",
      name: "Desenvolvimento de software",
      date: "2022-12-31T00:00:00.000Z",
      organization: "Cubos",
      location: "São Paulo",
    },
    {
      id: "1",
      name: "Desenvolvimento de software",
      date: "2022-12-31T00:00:00.000Z",
      organization: "Cubos",
      location: "São Paulo",
    },
    {
      id: "1",
      name: "Desenvolvimento de software",
      date: "2022-12-31T00:00:00.000Z",
      organization: "Cubos",
      location: "São Paulo",
    },
    {
      id: "1",
      name: "Desenvolvimento de software",
      date: "2022-12-31T00:00:00.000Z",
      organization: "Cubos",
      location: "São Paulo",
    },
  ]
  return (
    <main className="mt-10 flex flex-col">
      <Title>Eventos disponíveis</Title>
      <div className="mt-8 sm:grid sm:grid-cols-auto-fit-cards flex flex-wrap justify-center gap-x-2 gap-y-4">
        {events.map((event) => <EventCard key={event.id} event={event}></EventCard>)}
      </div>
    </main>
  )
}
