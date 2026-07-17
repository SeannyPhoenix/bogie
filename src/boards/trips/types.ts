import { RouteType } from "../departures/types";

export type RouteStopTimes = {
  id: string;
  name: string;
  agency: string;
  color: string;
  textColor: string;
  type: RouteType;
  stops: Stop[];
};

export type Stop = {
  id: string;
  name: string;
  times: Array<StopTime | null>;
};

export type StopTime = {
  tripId: string;
  arrival: Date;
  departure: Date;
  sequence: number;
  headsign: string;
};
