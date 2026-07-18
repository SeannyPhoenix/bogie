import { JSX } from "#bogie/jsx-runtime";

export type Status = "scheduled" | "on-time" | "early" | "late" | "canceled";

export const statusSymbols: Record<Status, string> = {
  scheduled: "\u25c7",
  "on-time": "\u25be",
  early: "\u25c2",
  late: "\u25b8",
  canceled: "",
};

export function StatusLegend(): JSX.Element {
  return (
    <div class="status-symbols">
      {Object.entries(statusSymbols).map(([status, symbol]) => (
        symbol !== "" ? (
        <div key={status}>
          {symbol}: {status}
        </div>
        ) : null
      ))}
    </div>
  );
}