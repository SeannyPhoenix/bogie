import { DepartureBoardContext } from "../departures/context";
import { renderRoute } from "../departures/route";
import { isParent, Stop, ParentStop } from "./types";

export function renderBoard(ctx: DepartureBoardContext): HTMLElement {
  const element = document.createElement("div");
  element.className = "board";

  const { now } = ctx;
  for (const stop of ctx.data.stops) {
    element.append(...renderStop(now, stop, false));
  }

  return element;
}

// There is only ever one level of nesting, so
// renderStop will never recurse more than once.
function renderStop(
  now: Date,
  stop: Stop | ParentStop,
  isChild: boolean,
): HTMLElement[] {
  const elements: HTMLElement[] = [];

  elements.push(renderStopName(stop.name, isChild));

  if (isParent(stop)) {
    for (const child of stop.children) {
      elements.push(...renderStop(now, child, true));
    }
    return elements;
  }

  for (const route of stop.routes) {
    elements.push(...renderRoute(now, route, isChild));
  }

  return elements;
}

function renderStopName(name: string, isChild: boolean): HTMLDivElement {
  const element = document.createElement("div");
  element.className = "stop-name";
  if (isChild) {
    element.classList.add("child-stop");
  }
  element.textContent = name;

  return element;
}
