import { RenderContext } from "./context";
import { timeAgo } from "./time";

export function renderHeader(ctx: RenderContext): HTMLDivElement {
  const element = document.createElement("div");
  element.className = "heading";

  element.append(renderUpcoming());
  element.append(renderUpdatedAt(ctx));

  return element;
}

function renderUpcoming(): HTMLDivElement {
  const element = document.createElement("div");
  element.className = "upcoming";
  element.textContent = "Upcoming Departures";
  return element;
}

function renderUpdatedAt(ctx: RenderContext): HTMLDivElement {
  const element = document.createElement("div");
  element.className = "updated-at";
  element.textContent = `Updated ${timeAgo(ctx)}`;
  return element;
}
