import { JSX } from "#bogie/jsx-runtime";
import { Departure } from "./departure";
import { Route, RouteType } from "./types";

const routeIcons: Record<RouteType, string> = {
  "light-rail": "🚈",
  subway: "🚇",
  rail: "🚆",
  bus: "🚌",
  ferry: "⛴️",
  "cable-tram": "🚡",
  "aerial-lift": "🚠",
  funicular: "🚞",
  trolleybus: "🚎",
  monorail: "🚝",
};

type RouteProps = {
  now: Date;
  route: Route;
  isChild: boolean;
};

export function Route({now, route, isChild}: RouteProps): JSX.Element {
  return (
    <>
      <RouteInfo route={route} isChild={isChild} />
      {route.departures.map((departure) => <Departure now={now} departure={departure} />)}
    </>
  );
}

type RouteInfoProps = {
  route: Route;
  isChild: boolean;
};

export function RouteInfo({route, isChild}: RouteInfoProps): JSX.Element {
  return (
    <div class={`route-info${isChild ? " child-stop" : ""}`}>
      <RouteIcon route={route} />
      <Headsign route={route} />
    </div>
  );
}

export function RouteIcon({route}:{route: Route}): JSX.Element {
  return (
    <div
      class="route-icon"
      style={`color:#${route.textColor};background-color:#${route.color};`}
    >
      {routeIcons[route.type]}
    </div>
  );
}

function Headsign({route}:{route: Route}): JSX.Element {
  return (
    <div class="headsign">
      {route.departures.length === 0 ? route.name : route.departures[0].headsign}
    </div>
  );
}
