import { StopDepartures } from "./types";

export type RenderContext = {
  now: Date;
  data: StopDepartures;
};

export function newRenderContext(data: StopDepartures): RenderContext {
  return {
    now: new Date(),
    data,
  };
}
