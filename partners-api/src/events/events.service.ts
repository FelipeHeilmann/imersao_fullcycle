import { Injectable } from '@nestjs/common';
import { CreateEventDto } from './dto/create-event.dto';
import { UpdateEventDto } from './dto/update-event.dto';
import { PrismaService } from 'src/prisma/prisma.service';
import { Prisma, SpotStatus, TicketStatus } from '@prisma/client';
import { ResertveSpotDto } from 'src/spots/dto/reserve-spot.dto';

@Injectable()
export class EventsService {
  constructor(private prismaService: PrismaService) {}

  async create(createEventDto: CreateEventDto) {
    return await this.prismaService.event.create({
      data: {
        ...createEventDto,
        date: new Date(createEventDto.date),
      },
    });
  }

  async findAll() {
    return await this.prismaService.event.findMany();
  }

  async findOne(id: string) {
    return await this.prismaService.event.findUnique({
      where: {
        id: id,
      },
    });
  }

  async update(id: string, updateEventDto: UpdateEventDto) {
    return await this.prismaService.event.update({
      data: {
        ...updateEventDto,
        date: new Date(updateEventDto.date),
      },
      where: { id: id },
    });
  }

  async remove(id: string) {
    return await this.prismaService.event.delete({
      where: { id: id },
    });
  }

  async reserveSpot(reserveSpotDto: ResertveSpotDto & { eventId: string }) {
    const spots = await this.prismaService.spot.findMany({
      where: {
        eventId: reserveSpotDto.eventId,
        name: {
          in: reserveSpotDto.spots,
        },
      },
    });

    if (spots.length !== reserveSpotDto.spots.length) {
      const foundSpotsName = spots.map((spot) => spot.name);
      const notFoundSpotsName = reserveSpotDto.spots.filter(
        (spotName) => !foundSpotsName.includes(spotName),
      );
      throw new Error(`${notFoundSpotsName.join(', ')} not found`);
    }

    try {
      const tickets = await this.prismaService.$transaction(
        async (prisma) => {
          await prisma.reservationHistory.createMany({
            data: spots.map((spot) => ({
              spotId: spot.id,
              ticketKind: reserveSpotDto.ticketKind,
              email: reserveSpotDto.email,
              status: TicketStatus.reserved,
            })),
          });

          await prisma.spot.updateMany({
            where: {
              id: {
                in: spots.map((spot) => spot.id),
              },
            },
            data: {
              status: SpotStatus.reserved,
            },
          });

          const tickets = await Promise.all(
            spots.map((spot) =>
              prisma.ticket.create({
                data: {
                  spotId: spot.id,
                  email: reserveSpotDto.email,
                  ticketKind: reserveSpotDto.ticketKind,
                },
              }),
            ),
          );

          return tickets;
        },
        { isolationLevel: Prisma.TransactionIsolationLevel.ReadCommitted },
      );
      return tickets;
    } catch (e) {
      if (e instanceof Prisma.PrismaClientKnownRequestError) {
        switch (e.code) {
          case 'P2002':
          case 'P2014':
            throw new Error('Some spots are already reserved');
        }
      }
      throw e;
    }
  }
}
