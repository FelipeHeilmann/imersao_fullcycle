import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
} from '@nestjs/common';
import { SpotsService } from './spots.service';
import { CreateSpotDto } from './dto/create-spot.dto';
import { UpdateSpotDto } from './dto/update-spot.dto';

@Controller('events/:eventId/spots')
export class SpotsController {
  constructor(private readonly spotsService: SpotsService) {}

  @Post()
  async create(
    @Body() createSpotDto: CreateSpotDto,
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
    @Body() updateSpotDto: UpdateSpotDto,
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
