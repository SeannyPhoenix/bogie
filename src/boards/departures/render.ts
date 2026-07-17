import { renderBoard } from "./board";
import { DepartureBoardContext } from "./context";
import { renderFooter } from "./footer";
import { renderHeader } from "./header";

export function renderDepartures(
  portal: HTMLElement,
  ctx: DepartureBoardContext,
): void {
  const display = document.createElement("div");
  display.className = "display";

  display.append(renderHeader(ctx));
  display.append(renderBoard(ctx));
  display.append(renderFooter(ctx));

  portal.replaceChildren(display);
}
