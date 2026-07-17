import { getMockData } from "./mockData";
import { StopDepartures } from "./types";

export type DepartureBoardContext = {
  now: Date;
  data: StopDepartures;
};

export function newDepartureBoardContext(): DepartureBoardContext {
  return {
    now: new Date(),
    data: getMockData(),
  };
}
