import { formats } from "../departures/time";
import { Stop, StopTime } from "./types";
import { renderPoint } from "./point";

export function renderStop(
  stop: Stop,
  index: number,
  stops: Stop[],
): HTMLElement[] {
  const elements: HTMLElement[] = [];

  const stopName = renderStopName(stop);
  elements.push(stopName);

  for (const time of stop.times) {
    const stopTimeElements = renderStopTimes(time, index, stops.length);
    elements.push(...stopTimeElements);
  }

  return elements;
}

function renderStopName(stop: Stop): HTMLElement {
  const stopName = document.createElement("div");
  stopName.classList.add("stop-name");
  stopName.innerText = stop.name;

  return stopName;
}

function renderStopTimes(
  stopTime: StopTime | null,
  index: number,
  length: number,
): HTMLElement[] {
  if (stopTime === null) {
    return [document.createElement("div"), document.createElement("div")];
  }
  const elements: HTMLElement[] = [];

  const point = renderPoint(index, length, stopTime.departure);
  elements.push(point);

  const departureTime = renderDepartureTime(stopTime);
  elements.push(departureTime);

  return elements;
}

function renderDepartureTime(stopTime: StopTime): HTMLElement {
  const departureTime = document.createElement("div");
  departureTime.classList.add("stop-time");
  departureTime.innerText = formats.timeOnly.format(stopTime.departure);

  return departureTime;
}
