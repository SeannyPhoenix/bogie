import { formats } from "../departures/time";
import { Stop, StopTime } from "./types";
import { Point } from "./point";
import { JSX } from "#bogie/jsx-runtime";

export function Stop(
  stop: Stop,
  index: number,
  stops: Stop[],
): JSX.Element {
  return (
    <>
      <div class="trip-stop-name" key={stop.id}>{stop.name}</div>
      {stop.times.map((time) => (
          <StopTimes 
            stopTime={time}
            index={index}
            length={stops.length}
          />
      ))}
    </>
  );
}

type StopTimesProps = {
  stopTime: StopTime | null;
  index: number;
  length: number;
};

function StopTimes({stopTime, index, length}: StopTimesProps): JSX.Element {
  if (stopTime === null) {
    return <><div/><div/></>;
  }

  return (
    <>
      <Point
        current={index}
        total={length}
        time={stopTime.departure} 
      />
      <div class="stop-time" >
        {formats.timeOnly.format(stopTime.departure)}
      </div>
    </>
  );
}
