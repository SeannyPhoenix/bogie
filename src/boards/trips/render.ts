import "./trip.css";
import { TripBoardContext } from "./context";
import { renderStop } from "./stop";

export function renderTrip(portal: HTMLElement, ctx: TripBoardContext) {
  const element = document.createElement("div");

  const routeStopTimesElement = renderRouteStopTimes(ctx);
  element.append(routeStopTimesElement);

  portal.replaceChildren(element);
}

function renderRouteStopTimes(ctx: TripBoardContext): HTMLElement {
  const element = document.createElement("div");
  element.classList.add("trip-board");

  const route = ctx.data;
  for (const stop of route.stops.map(renderStop)) {
    element.append(...stop);
  }

  return element;
}
