import { Injectable } from '@nestjs/common';
import { CreateSpotDto } from './dto/create-spot.dto';
import { UpdateSpotDto } from './dto/update-spot.dto';
import { SpotStatus } from '@prisma/client';
import { PrismaService } from '../prisma/prisma.service';

@Injectable()
export class SpotsService {
  constructor(private prismaService: PrismaService) {}

  async create(createSpotDto: CreateSpotDto & { eventId: string }) {
    const event = await this.prismaService.event.findUnique({
      where: {
        id: createSpotDto.eventId,
      },
    });

    if (!event) throw new Error('Event not found');

    return await this.prismaService.spot.create({
      data: {
        ...createSpotDto,
        status: SpotStatus.available,
      },
    });
  }

  async findAll(eventId: string) {
    return await this.prismaService.spot.findMany({
      where: {
        eventId,
      },
    });
  }

  async findOne(eventId: string, spotId: string) {
    return await this.prismaService.spot.findFirst({
      where: {
        id: spotId,
        eventId,
      },
    });
  }

  async update(eventId: string, spotId: string, updateSpotDto: UpdateSpotDto) {
    return await this.prismaService.spot.update({
      where: {
        id: spotId,
        eventId,
      },
      data: updateSpotDto,
    });
  }

  async remove(eventId: string, spotId: string) {
    return await this.prismaService.spot.delete({
      where: {
        id: spotId,
        eventId,
      },
    });
  }
}
