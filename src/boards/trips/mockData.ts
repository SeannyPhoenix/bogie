import { dateAfterMinutes } from "../departures/mockData";
import { Trip } from "./types";

export function getMockTripData(): Trip {
  const now = new Date();

  return {
    id: "BA:1951767",
    route: {
      id: "BA:Orange-N",
      name: "Orange-N",
      color: "FF9933",
      textColor: "000000",
      type: "subway",
      agency: "BART",
    },
    stopTimes: [
      {
        sequence: 0,
        stop: {
          id: "BA:12TH",
          name: "12th St. Oakland City Center",
        },
        headsign: "Richmond",
        arrival: dateAfterMinutes(now, 4),
        departure: dateAfterMinutes(now, 4),
      },
    ],
  };
}
