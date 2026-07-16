import { renderBoard } from "./board";
import { RenderContext } from "./context";
import { renderFooter } from "./footer";
import { renderHeader } from "./header";

export function render(portal: HTMLElement, ctx: RenderContext): void {
  const display = document.createElement("div");
  display.className = "display";

  display.append(renderHeader(ctx));
  display.append(renderBoard(ctx));
  display.append(renderFooter(ctx));

  portal.replaceChildren(display);
}
