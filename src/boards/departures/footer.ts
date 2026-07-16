import { RenderContext } from "./context";
import { statusSymbols } from "./symbols";
import { timeFetched } from "./time";

export function renderFooter(ctx: RenderContext): HTMLDivElement {
  const element = document.createElement("div");
  element.className = "footer";
  element.append(renderStatusSymbols());
  element.append(renderFetchedAt(ctx));
  return element;
}

function renderStatusSymbols(): HTMLDivElement {
  const element = document.createElement("div");
  element.className = "status-symbols";

  for (const status of [
    "scheduled",
    "on-time",
    "early",
    "late",
    // "canceled",
  ] as const) {
    const symbolElement = document.createElement("div");
    symbolElement.textContent = `${statusSymbols[status]}: ${status}`;
    element.append(symbolElement);
  }

  return element;
}

function renderFetchedAt(ctx: RenderContext): HTMLDivElement {
  const element = document.createElement("div");
  element.className = "fetched-at";
  element.textContent = `Fetched at ${timeFetched(ctx)}`;
  return element;
}
