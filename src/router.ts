import { newRenderContext } from "./boards/departures/context";
import { getMockData } from "./boards/departures/mockData";
import { render } from "./boards/departures/render";
import { newTripsBoardContext } from "./boards/trips/context";
import { getMockTripData } from "./boards/trips/mockData";
import { renderTrip } from "./boards/trips/render";

type RenderFunction = () => void;

export function route(): RenderFunction {
  const { pathname, search } = window.location;

  const route = routes[pathname];
  if (route) {
    return route;
  }

  console.warn(`No route found for pathname: ${pathname}`);
  return () => {
    console.log(`Pathname: ${pathname}, Search: ${search}`);
  };
}

const routes: Record<string, RenderFunction> = {
  "/": () => {
    console.log("Rendering home page");
  },
  "/departures": () => {
    const mockData = getMockData();
    const { body } = document;

    function update() {
      const ctx = newRenderContext(mockData);
      render(body, ctx);
    }
    update();

    // setInterval(update, 1000);
  },
  "/trips": () => {
    const mockData = getMockTripData();
    const { body } = document;

    function update() {
      const ctx = newTripsBoardContext(mockData);
      renderTrip(body, ctx);
    }
    update();
  },
};
