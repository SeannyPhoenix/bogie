import { renderBoard } from "./board";
import { DepartureBoardContext } from "./context";
import { renderFooter } from "./footer";
import { renderHeader } from "./header";

export function renderDepartures(
  portal: HTMLElement,
  ctx: DepartureBoardContext,
): void {
  const content = (
    <div class="display">
      {renderHeader(ctx)}
      {renderBoard(ctx)}
      {renderFooter(ctx)}
    </div>
  );

  portal.replaceChildren(content);
}
