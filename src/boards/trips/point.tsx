import { JSX } from "#bogie/jsx-runtime";
import { formats } from "../departures/time";

type PointType = "start" | "middle" | "end";

function pointType(current: number, total: number): PointType {
  switch (current) {
    case 0:
      return "start";
    case total - 1:
      return "end";
    default:
      return "middle";
  }
}
 
type StopStatus = "complete" | "current" | "incomplete";

function pointStatus(time: Date): StopStatus {
  const now = new Date();
  const lastMinute = new Date(now);
  lastMinute.setSeconds(now.getSeconds() - 15);
  const nextMinute = new Date(now);
  nextMinute.setSeconds(now.getSeconds() + 15);

  if (time.getTime() < lastMinute.getTime()) {
    return "complete";
  }

  if (time.getTime() > nextMinute.getTime()) {
    return "incomplete";
  }

  return "current";
}

type PointProps = {
  current: number;
  total: number;
  time: Date;
};

export function Point({current, total, time}: PointProps): JSX.Element {
  const type = pointType(current, total);
  const status = pointStatus(time);

  const outerPoint = <div class={`point ${status} incomplete`}/>;
  const innerPoint = status === "current" ? <div class="inner-point complete"/> : null;

  const upperBitStyle = status === "incomplete" ? "incomplete" : "complete";
  const upperBit = type === "middle" || type === "end" ? <div class={`upper-bit ${upperBitStyle}`}/> : null;

  const lowerBitStatus = status === "complete" ? "complete" : "incomplete";
  const lowerBit = type === "start" || type === "middle" ? <div class={`lower-bit ${lowerBitStatus}`}/> : null;

  return (
    <>
      <div class="trip-symbol">
        {outerPoint}
        {innerPoint}
        {upperBit}
        {lowerBit}
      </div>
    </>
  )
}
