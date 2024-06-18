import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
} from '@nestjs/common';
import { SpotsService } from '@app/core/spots/spots.service';
import { CreateSpotRequest } from './request/create-spot.request';
import { UpdateSpotRequest } from './request/update-spot.request';

@Controller('events/:eventId/spots')
export class SpotsController {
  constructor(private readonly spotsService: SpotsService) {}

  @Post()
  async create(
    @Body() createSpotDto: CreateSpotRequest,
    @Param('eventId') eventId: string,
  ) {
    return await this.spotsService.create({
      ...createSpotDto,
      eventId,
    });
  }

  @Get()
  async findAll(@Param('eventId') eventId: string) {
    return await this.spotsService.findAll(eventId);
  }

  @Get(':spotId')
  async findOne(
    @Param('spotId') spotId: string,
    @Param('eventId') eventId: string,
  ) {
    return await this.spotsService.findOne(eventId, spotId);
  }

  @Patch(':eventId')
  async update(
    @Param('spotId') spotId: string,
    @Param('eventId') eventId: string,
    @Body() updateSpotDto: UpdateSpotRequest,
  ) {
    return await this.spotsService.update(eventId, spotId, updateSpotDto);
  }

  @Delete(':spotId')
  async remove(
    @Param('spotId') spotId: string,
    @Param('eventId') eventId: string,
  ) {
    return await this.spotsService.remove(eventId, spotId);
  }
}
