import { routeTypeMap } from "../departures/types";
import freshData from "./fresh.json";
import { RouteStopTimes } from "./types";

export function getFreshData(): RouteStopTimes {
  const first = freshData[0];
  const routeStopTimes: RouteStopTimes = {
    id: first.routeId,
    name: first.routeName,
    agency: first.agencyId,
    color: first.routeColor,
    textColor: first.routeTextColor,
    type: routeTypeMap[first.routeType],
    stops: [],
  };

  for (const row of freshData) {
    let stopIndex = findStop(routeStopTimes, row.stopId);
    if (stopIndex === undefined) {
      const stop = {
        id: row.stopId,
        name: row.stopName,
        times: [],
      };
      routeStopTimes.stops.push(stop);
      stopIndex = routeStopTimes.stops.length - 1;
    }

    const stop = routeStopTimes.stops[stopIndex];
    const stopTime = {
      tripId: row.tripId,
      arrival: parseDate(row.arrivalTime),
      departure: parseDate(row.departureTime),
      sequence: row.stopSequence,
      headsign: row.stopHeadsign,
    };
    stop.times.push(stopTime);
    routeStopTimes.stops[stopIndex] = stop;
  }

  return routeStopTimes;
}

function parseDate(dateString: string): Date {
  const date = new Date();
  date.setHours(parseInt(dateString.slice(0, 2), 10));
  date.setMinutes(parseInt(dateString.slice(3, 5), 10));
  date.setSeconds(parseInt(dateString.slice(6, 8), 10));
  return date;
}

function findStop(rst: RouteStopTimes, stopId: string): number | undefined {
  for (let i = 0; i < rst.stops.length; i++) {
    if (rst.stops[i].id === stopId) {
      return i;
    }
  }
  return undefined;
}
