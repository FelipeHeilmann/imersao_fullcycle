import { TicketKind } from '@prisma/client';

export class ResertveSpotDto {
  spots: string[];
  ticketKind: TicketKind;
  email: string;
}
