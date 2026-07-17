import "./trip.css";
import { TripsBoardContext } from "./context";
import { renderStop } from "./stop";

export function renderTrip(portal: HTMLElement, ctx: TripsBoardContext) {
  const element = document.createElement("div");

  const routeStopTimesElement = renderRouteStopTimes(ctx);
  element.append(routeStopTimesElement);

  portal.replaceChildren(element);
}

function renderRouteStopTimes(ctx: TripsBoardContext): HTMLElement {
  const element = document.createElement("div");
  element.classList.add("trip-board");

  const route = ctx.data;
  for (const stop of route.stops.map(renderStop)) {
    element.append(...stop);
  }

  return element;
}
