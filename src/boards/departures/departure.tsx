import { JSX } from "#bogie/jsx-runtime";
import { Status, statusSymbols } from "./symbols";
import { Departure as DepartureModel } from "./types";

type DepartureProps = {
  now: Date;
  departure: DepartureModel;
};

export function Departure({now, departure}: DepartureProps): JSX.Element {
  const departureStatus = status(departure);

  return (
    <div class={`departure ${departureStatus}`}>
      <a class="no-link" href={link(departure)}>
        {formatTime(now, departure)} {statusSymbols[departureStatus]}
      </a>
    </div>
  );
}

function status(departure: DepartureModel): Status {
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

function formatTime(now: Date, departure: DepartureModel): string {
  const time = departure.predicted ?? departure.scheduled;
  const diffMinutes = Math.floor((time.getTime() - now.getTime()) / 60_000);

  if (diffMinutes < 1) {
    return "now";
  }

  return `${diffMinutes} min`;
}

function link(departure: DepartureModel): string {
  const { tripId } = departure;
  return `/trips/${tripId}`;
}