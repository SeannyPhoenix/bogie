import { Status, statusSymbols } from "./symbols";
import { Departure } from "./types";

export function renderDeparture(now: Date, departure: Departure): HTMLElement {
  const element = document.createElement("div");
  const departureStatus = status(departure);

  element.className = `departure ${departureStatus}`;
  element.textContent = `${formatTime(now, departure)} ${statusSymbols[departureStatus]}`;

  return element;
}

function status(departure: Departure): Status {
  if (departure.canceled) {
    return "canceled";
  }

  const { scheduled, predicted } = departure;

  if (predicted) {
    const diff = Math.floor(
      (predicted.getTime() - scheduled.getTime()) / 60_000,
    );
    if (diff < 0) {
      return "early";
    }
    if (diff > 0) {
      return "late";
    }

    return "on-time";
  }

  return "scheduled";
}

function formatTime(now: Date, departure: Departure): string {
  const time = departure.predicted ?? departure.scheduled;
  const diffMinutes = Math.floor((time.getTime() - now.getTime()) / 60_000);

  if (diffMinutes < 1) {
    return "now";
  }

  return `${diffMinutes} min`;
}
