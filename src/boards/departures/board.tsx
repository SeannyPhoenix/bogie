import { JSX } from "#bogie/jsx-runtime";
import { DepartureBoardContext } from "./context";
import { Route } from "./route";
import { isParent, Stop, ParentStop } from "./types";

export function renderBoard({now, data}: DepartureBoardContext): JSX.Element {
  return (
    <div class="board">
      {data.stops.map((stop) => <Stop now={now} stop={stop} />)}
    </div>
  );
}

type StopProps = {
  now: Date;
  stop: Stop | ParentStop;
  isChild?: boolean;
};
// There is only ever one level of nesting, so
// renderStop will never recurse more than once.
function Stop({now, stop, isChild = false}: StopProps): JSX.Element {
  const stops = isParent(stop) 
    ? stop.children.map((stop) => <Stop now={now} stop={stop} isChild={true} />) 
    : stop.routes.map((route) => <Route now={now} route={route} isChild={isChild} />)

  return (
    <>
      <StopName name={stop.name} isChild={isChild} />
      {stops}
    </>
  );
}

type StopNameProps = {
  name: string;
  isChild: boolean;
};

function StopName({name, isChild}: StopNameProps): JSX.Element {
  return (
    <div class={`stop-name${isChild ? " child-stop" : ""}`}>
      {name}
    </div>
  );
}
