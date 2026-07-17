export type StopId = string;

export type StopDepartures = {
  asOf: Date;
  stops: Array<Stop | ParentStop>;
};

export type ParentStop = {
  id: string;
  name: string;
  agency: string;
  children: Stop[];
};

export type Stop = {
  id: string;
  name: string;
  agency: string;
  routes: Route[];
};

export type Route = {
  id: string;
  name: string;
  color: string;
  textColor: string;
  type: RouteType;
  // direction: string;
  departures: Departure[];
};

export type Departure = {
  tripId: string;
  headsign: string;
  scheduled: Date;
} & (
  | {
      canceled: true;
      predicted: null;
    }
  | {
      canceled: false;
      predicted: Date | null;
    }
);

export type RouteType =
  | "light-rail"
  | "subway"
  | "rail"
  | "bus"
  | "ferry"
  | "cable-tram"
  | "aerial-lift"
  | "funicular"
  | "trolleybus"
  | "monorail";

export const routeTypeMap: Record<number, RouteType> = {
  0: "light-rail",
  1: "subway",
  3: "bus",
  2: "rail",
  4: "ferry",
  5: "cable-tram",
  6: "aerial-lift",
  7: "funicular",
  11: "trolleybus",
  12: "monorail",
};

export function isParent(stop: Stop | ParentStop): stop is ParentStop {
  return "children" in stop;
}
