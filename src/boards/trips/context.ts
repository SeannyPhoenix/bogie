import { Trip } from "./types";

export type TripsBoardContext = {
  now: Date;
  data: Trip;
};

export function newTripsBoardContext(data: Trip): TripsBoardContext {
  return {
    now: new Date(),
    data,
  };
}
