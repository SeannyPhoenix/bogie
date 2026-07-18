import { JSX } from "#bogie/jsx-runtime";
import { DepartureBoardContext } from "./context";
import { StatusLegend, statusSymbols } from "./symbols";
import { timeFetched } from "./time";


export function renderFooter(ctx: DepartureBoardContext): JSX.Element {
  return (
    <div class="footer">
      <StatusLegend/>
      <div class="fetched-at">
        Fetched at {timeFetched(ctx)}
      </div>
    </div>
  );
}
