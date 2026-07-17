import { DepartureBoardContext } from "./context";

export const formats = {
  timeOnly: new Intl.DateTimeFormat("en-US", {
    hour: "2-digit",
    minute: "2-digit",
    hour12: false,
  }),
};

export function timeAgo(ctx: DepartureBoardContext): string {
  const {
    now,
    data: { asOf },
  } = ctx;
  const diffMs = now.getTime() - asOf.getTime();
  const diffSeconds = Math.floor(diffMs / 1_000);

  if (diffSeconds < 10) {
    return "just now";
  }

  if (diffSeconds < 60) {
    return "less than a minute ago";
  }

  const diffMinutes = Math.floor(diffMs / 60_000);

  if (diffMinutes === 1) {
    return "1 minute ago";
  }

  return `${diffMinutes} minutes ago`;
}

export function timeFetched(ctx: DepartureBoardContext): string {
  const { asOf } = ctx.data;
  return formats.timeOnly.format(asOf);
}
