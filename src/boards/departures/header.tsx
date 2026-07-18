import { JSX } from "#bogie/jsx-runtime";
import { DepartureBoardContext } from "./context";
import { timeAgo } from "./time";

export function renderHeader(ctx: DepartureBoardContext): JSX.Element {
  return (
    <div class="heading">
      <div class="upcoming">Upcoming Departures</div>
      <div class="updated-at">Updated {timeAgo(ctx)}</div>
    </div>
  );
}
