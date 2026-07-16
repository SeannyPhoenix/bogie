import { StopDepartures } from "./types";

export function getMockData(): StopDepartures {
  const now = new Date();

  const unified: StopDepartures = {
    asOf: now,
    stops: [
      {
        id: "mtc:12th-st-oakland-city-center-bart",
        name: "12th Street / Oakland City Center",
        agency: "BART",
        routes: [
          {
            id: "BA:Orange-N",
            name: "Orange-N",
            color: "FF9933",
            textColor: "000000",
            type: "subway",
            departures: [
              {
                tripId: "BA:1951767",
                headsign: "Richmond",
                scheduled: dateAfterMinutes(now, 2),
                canceled: false,
                predicted: dateAfterMinutes(now, 2),
              },
              {
                tripId: "BA:1951768",
                headsign: "Richmond",
                scheduled: dateAfterMinutes(now, 13),
                canceled: false,
                predicted: dateAfterMinutes(now, 15),
              },
            ],
          },
          {
            id: "BA:Red-N",
            name: "Red-N",
            color: "FF0000",
            textColor: "FFFFFF",
            type: "subway",
            departures: [
              {
                tripId: "BA:1951769",
                headsign: "Richmond",
                scheduled: dateAfterMinutes(now, 3),
                canceled: false,
                predicted: dateAfterMinutes(now, 5),
              },
              {
                tripId: "BA:1951770",
                headsign: "Richmond",
                scheduled: dateAfterMinutes(now, 12),
                canceled: false,
                predicted: dateAfterMinutes(now, 12),
              },
              {
                tripId: "BA:1951778",
                headsign: "Richmond",
                scheduled: dateAfterMinutes(now, 22),
                canceled: false,
                predicted: dateAfterMinutes(now, 22),
              },
            ],
          },

          {
            id: "BA:Red-S",
            name: "Red-S",
            color: "FF0000",
            textColor: "FFFFFF",
            type: "subway",
            departures: [
              {
                tripId: "BA:1951771",
                headsign: "Millbrae",
                scheduled: dateAfterMinutes(now, 4),
                canceled: true,
                predicted: null,
              },
              {
                tripId: "BA:1951779",
                headsign: "Millbrae",
                scheduled: dateAfterMinutes(now, 15),
                canceled: false,
                predicted: dateAfterMinutes(now, 14),
              },
              {
                tripId: "BA:1951780",
                headsign: "Millbrae",
                scheduled: dateAfterMinutes(now, 25),
                canceled: false,
                predicted: dateAfterMinutes(now, 25),
              },
            ],
          },
          {
            id: "BA:Yellow-S",
            name: "Yellow-S",
            color: "FFFF33",
            textColor: "000000",
            type: "subway",
            departures: [
              {
                tripId: "BA:1951781",
                headsign: "San Francisco International Airport",
                scheduled: dateAfterMinutes(now, 6),
                canceled: false,
                predicted: dateAfterMinutes(now, 6),
              },
              {
                tripId: "BA:1951774",
                headsign: "San Francisco International Airport",
                scheduled: dateAfterMinutes(now, 16),
                canceled: false,
                predicted: null,
              },
              {
                tripId: "BA:1951782",
                headsign: "San Francisco International Airport",
                scheduled: dateAfterMinutes(now, 26),
                canceled: false,
                predicted: dateAfterMinutes(now, 27),
              },
            ],
          },
          {
            id: "BA:Orange-S",
            name: "Orange-S",
            color: "FF9933",
            textColor: "000000",
            type: "subway",
            departures: [
              {
                tripId: "BA:1951783",
                headsign: "Berryessa/North San Jose",
                scheduled: dateAfterMinutes(now, 7),
                canceled: false,
                predicted: dateAfterMinutes(now, 7),
              },
              {
                tripId: "BA:1951776",
                headsign: "Berryessa/North San Jose",
                scheduled: dateAfterMinutes(now, 17),
                canceled: false,
                predicted: dateAfterMinutes(now, 20),
              },
              {
                tripId: "BA:1951777",
                headsign: "Berryessa/North San Jose",
                scheduled: dateAfterMinutes(now, 25),
                canceled: false,
                predicted: dateAfterMinutes(now, 25),
              },
            ],
          },

          {
            id: "BA:Yellow-N",
            name: "Yellow-N",
            color: "FFFF33",
            textColor: "000000",
            type: "subway",
            departures: [
              {
                tripId: "BA:1951784",
                headsign: "SFO / Pittsburg/Bay Point",
                scheduled: dateAfterMinutes(now, 10),
                canceled: false,
                predicted: dateAfterMinutes(now, 10),
              },
              {
                tripId: "BA:1951785",
                headsign: "SFO / SF / Antioch",
                scheduled: dateAfterMinutes(now, 20),
                canceled: false,
                predicted: dateAfterMinutes(now, 20),
              },
              {
                tripId: "BA:1951786",
                headsign: "SFO / SF / Antioch",
                scheduled: dateAfterMinutes(now, 30),
                canceled: false,
                predicted: dateAfterMinutes(now, 33),
              },
            ],
          },
        ],
      },
    ],
  };

  const separate: StopDepartures = {
    asOf: now,
    stops: [
      {
        id: "mtc:12th-st-oakland-city-center-bart",
        name: "12th Street / Oakland City Center",
        agency: "BART",
        children: [
          {
            id: "900101",
            name: "Platform 1",
            agency: "BART",
            routes: [
              {
                id: "BA:Orange-N",
                name: "Orange-N",
                color: "FF9933",
                textColor: "000000",
                type: "subway",
                departures: [
                  {
                    tripId: "BA:1951767",
                    headsign: "Richmond",
                    scheduled: dateAfterMinutes(now, 2),
                    canceled: false,
                    predicted: dateAfterMinutes(now, 2),
                  },
                  {
                    tripId: "BA:1951768",
                    headsign: "Richmond",
                    scheduled: dateAfterMinutes(now, 13),
                    canceled: false,
                    predicted: dateAfterMinutes(now, 15),
                  },
                ],
              },
              {
                id: "BA:Red-N",
                name: "Red-N",
                color: "FF0000",
                textColor: "FFFFFF",
                type: "subway",
                departures: [
                  {
                    tripId: "BA:1951769",
                    headsign: "Richmond",
                    scheduled: dateAfterMinutes(now, 3),
                    canceled: false,
                    predicted: dateAfterMinutes(now, 5),
                  },
                  {
                    tripId: "BA:1951770",
                    headsign: "Richmond",
                    scheduled: dateAfterMinutes(now, 12),
                    canceled: false,
                    predicted: dateAfterMinutes(now, 12),
                  },
                  {
                    tripId: "BA:1951778",
                    headsign: "Richmond",
                    scheduled: dateAfterMinutes(now, 22),
                    canceled: false,
                    predicted: dateAfterMinutes(now, 22),
                  },
                ],
              },
            ],
          },
          {
            id: "900102",
            name: "Platform 2",
            agency: "BART",
            routes: [
              {
                id: "BA:Red-S",
                name: "Red-S",
                color: "FF0000",
                textColor: "FFFFFF",
                type: "subway",
                departures: [
                  {
                    tripId: "BA:1951771",
                    headsign: "Millbrae",
                    scheduled: dateAfterMinutes(now, 4),
                    canceled: true,
                    predicted: null,
                  },
                  {
                    tripId: "BA:1951779",
                    headsign: "Millbrae",
                    scheduled: dateAfterMinutes(now, 15),
                    canceled: false,
                    predicted: dateAfterMinutes(now, 14),
                  },
                  {
                    tripId: "BA:1951780",
                    headsign: "Millbrae",
                    scheduled: dateAfterMinutes(now, 25),
                    canceled: false,
                    predicted: dateAfterMinutes(now, 25),
                  },
                ],
              },
              {
                id: "BA:Yellow-S",
                name: "Yellow-S",
                color: "FFFF33",
                textColor: "000000",
                type: "subway",
                departures: [
                  {
                    tripId: "BA:1951781",
                    headsign: "San Francisco International Airport",
                    scheduled: dateAfterMinutes(now, 6),
                    canceled: false,
                    predicted: dateAfterMinutes(now, 6),
                  },
                  {
                    tripId: "BA:1951774",
                    headsign: "San Francisco International Airport",
                    scheduled: dateAfterMinutes(now, 16),
                    canceled: false,
                    predicted: null,
                  },
                  {
                    tripId: "BA:1951782",
                    headsign: "San Francisco International Airport",
                    scheduled: dateAfterMinutes(now, 26),
                    canceled: false,
                    predicted: dateAfterMinutes(now, 27),
                  },
                ],
              },
              {
                id: "BA:Orange-S",
                name: "Orange-S",
                color: "FF9933",
                textColor: "000000",
                type: "subway",
                departures: [
                  {
                    tripId: "BA:1951783",
                    headsign: "Berryessa/North San Jose",
                    scheduled: dateAfterMinutes(now, 7),
                    canceled: false,
                    predicted: dateAfterMinutes(now, 7),
                  },
                  {
                    tripId: "BA:1951776",
                    headsign: "Berryessa/North San Jose",
                    scheduled: dateAfterMinutes(now, 17),
                    canceled: false,
                    predicted: dateAfterMinutes(now, 20),
                  },
                  {
                    tripId: "BA:1951777",
                    headsign: "Berryessa/North San Jose",
                    scheduled: dateAfterMinutes(now, 25),
                    canceled: false,
                    predicted: dateAfterMinutes(now, 25),
                  },
                ],
              },
            ],
          },
          {
            id: "900103",
            name: "Platform 3",
            agency: "BART",
            routes: [
              {
                id: "BA:Yellow-N",
                name: "Yellow-N",
                color: "FFFF33",
                textColor: "000000",
                type: "subway",
                departures: [
                  {
                    tripId: "BA:1951784",
                    headsign: "SFO / Pittsburg/Bay Point",
                    scheduled: dateAfterMinutes(now, 10),
                    canceled: false,
                    predicted: dateAfterMinutes(now, 10),
                  },
                  {
                    tripId: "BA:1951785",
                    headsign: "SFO / SF / Antioch",
                    scheduled: dateAfterMinutes(now, 20),
                    canceled: false,
                    predicted: dateAfterMinutes(now, 20),
                  },
                  {
                    tripId: "BA:1951786",
                    headsign: "SFO / SF / Antioch",
                    scheduled: dateAfterMinutes(now, 30),
                    canceled: false,
                    predicted: dateAfterMinutes(now, 33),
                  },
                ],
              },
            ],
          },
        ],
      },
    ],
  };

  return separate;
}

function dateAfterMinutes(time: Date, minutesFromNow: number) {
  const randomSeconds = Math.floor(Math.random() * 60);
  return new Date(
    time.getTime() + minutesFromNow * 60_000 + randomSeconds * 1_000,
  );
}
