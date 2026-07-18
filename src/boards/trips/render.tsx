import "./trip.css";
import { TripBoardContext } from "./context";
import { Stop } from "./stop";
import { JSX } from "#bogie/jsx-runtime";

export function renderTrip(portal: HTMLElement, {data: {stops}}: TripBoardContext) {
  const content = (
    <div class="trip-board">
      {stops.map(Stop)}
    </div>
  );

  portal.replaceChildren(content);
}
