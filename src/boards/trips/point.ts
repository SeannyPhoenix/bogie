type PointType = "start" | "middle" | "end";

function pointType(current: number, total: number): PointType {
  if (current === 0) {
    return "start";
  }

  if (current === total - 1) {
    return "end";
  }

  return "middle";
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

export function renderPoint(
  current: number,
  total: number,
  time: Date,
): HTMLElement {
  const element = document.createElement("div");
  element.classList.add("trip-symbol");

  const type = pointType(current, total);
  const status = pointStatus(time);

  const point = document.createElement("div");
  point.classList.add("point");
  point.classList.add(status);
  point.classList.add("incomplete");
  element.append(point);

  if (status === "current") {
    const innerPoint = document.createElement("div");
    innerPoint.classList.add("inner-point");
    innerPoint.classList.add("complete");
    element.append(innerPoint);
  }

  if (type === "middle" || type === "end") {
    const bit = document.createElement("div");
    bit.classList.add("upper-bit");
    switch (status) {
      case "complete":
        bit.classList.add("complete");
        break;
      case "current":
        bit.classList.add("complete");
        break;
      case "incomplete":
        bit.classList.add("incomplete");
        break;
    }
    element.append(bit);
  }

  if (type === "start" || type === "middle") {
    const bit = document.createElement("div");
    bit.classList.add("lower-bit");
    switch (status) {
      case "complete":
        bit.classList.add("complete");
        break;
      case "current":
        bit.classList.add("incomplete");
        break;
      case "incomplete":
        bit.classList.add("incomplete");
        break;
    }
    element.append(bit);
  }

  return element;
}
