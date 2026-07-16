import { TripsBoardContext } from "./context";

export function renderTrip(portal: HTMLElement, ctx: TripsBoardContext) {
  const element = document.createElement("div");
  element.innerText = `Trip ID: ${ctx.data.id}, Route: ${ctx.data.route.name}`;
  portal.replaceChildren(element);
}
