import { renderDeparture } from "./departure";
import { Route, RouteType } from "./types";

const routeIcons: Record<RouteType, string> = {
  "light-rail": "🚈",
  subway: "🚇",
  rail: "🚆",
  ferry: "⛴️",
  "cable-tram": "🚡",
  "aerial-lift": "🚠",
  funicular: "🚞",
  trolleybus: "🚎",
  monorail: "🚝",
};

export function renderRoute(
  now: Date,
  route: Route,
  isChild: boolean,
): HTMLElement[] {
  const elements: HTMLElement[] = [];

  elements.push(renderRouteInfo(route, isChild));

  for (const departure of route.departures) {
    elements.push(renderDeparture(now, departure));
  }

  return elements;
}

export function renderRouteInfo(route: Route, isChild: boolean): HTMLElement {
  const element = document.createElement("div");
  element.className = "route-info";
  if (isChild) {
    element.classList.add("child-stop");
  }

  element.appendChild(renderRouteIcon(route));
  element.appendChild(renderHeadsign(route));
  return element;
}

export function renderRouteIcon(route: Route): HTMLElement {
  const element = document.createElement("div");
  element.className = "route-icon";
  element.textContent = routeIcons[route.type];
  element.style.color = `#${route.textColor}`;
  element.style.backgroundColor = `#${route.color}`;

  return element;
}

function renderHeadsign(route: Route): HTMLElement {
  const element = document.createElement("div");
  element.className = "headsign";

  if (route.departures.length == 0) {
    element.textContent = route.name;
    return element;
  }

  element.textContent = route.departures[0].headsign;
  return element;
}
