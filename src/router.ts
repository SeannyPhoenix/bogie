import { newDepartureBoardContext } from "./boards/departures/context";
import { renderDepartures } from "./boards/departures/render";
import { newTripBoardContext } from "./boards/trips/context";
import { renderTrip } from "./boards/trips/render";

type RenderFunction = () => void;

export function route(): RenderFunction {
  const { pathname, search } = window.location;

  // get the first path segment
  const first = pathname.split("/")[1];

  const render = routes[first ?? ""];
  if (render) {
    return render;
  }

  return () => {
    console.warn(`No route found for pathname: ${pathname}`);
    console.log(`Pathname: ${pathname}, Search: ${search}`);
  };
}

const routes: Record<string, RenderFunction> = {
  "": () => {
    console.log("Rendering home page");
  },
  departures: () => {
    const ctx = newDepartureBoardContext();
    const { body } = document;
    renderDepartures(body, ctx);
  },
  trips: renderTripsRoute,
};

function renderTripsRoute() {
  const ctx = newTripBoardContext();
  const { body } = document;
  renderTrip(body, ctx);
}
