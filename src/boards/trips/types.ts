import { RouteType } from "../departures/types";

export type Trip = {
  id: string;
  route: Route;
  stopTimes: StopTime[];
};

export type Route = {
  id: string;
  name: string;
  agency: string;
  color: string;
  textColor: string;
  type: RouteType;
};

type StopTime = {
  stop: Stop;
  arrival: Date;
  departure: Date;
  sequence: number;
  headsign: string;
};

type Stop = {
  id: string;
  name: string;
};
