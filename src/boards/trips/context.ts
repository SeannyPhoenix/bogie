import { RouteStopTimes } from "./types";
import { getFreshData } from "./freshData";

export type TripsBoardContext = {
  now: Date;
  data: RouteStopTimes;
  completed?: number;
};

export function newTripsBoardContext(): TripsBoardContext {
  const ctx: TripsBoardContext = {
    now: new Date(),
    data: getFreshData(),
    completed: getCompleted(),
  };
  return ctx;
}

function getCompleted(): number | undefined {
  const { search } = window.location;
  const params = new URLSearchParams(search);

  if (params.has("completed")) {
    const c = params.get("completed");
    const completed = parseInt(c ?? "0", 10);
    if (isNaN(completed)) {
      return undefined;
    }
    return completed;
  }

  return undefined;
}
