"use client";

import { Forecast } from "@/types";

interface Props {
  data: Forecast | undefined;
  loading: boolean;
}

function formatDate(dateStr: string) {
  const date = new Date(dateStr);
  return date.toLocaleDateString("en-US", {
    weekday: "short",
    month: "short",
    day: "numeric",
  });
}

export default function ForecastStrip({ data, loading }: Props) {
  if (loading) {
    return (
      <div className="forecast-strip">
        {Array.from({ length: 5 }).map((_, i) => (
          <div key={i} className="forecast-day-card">
            <div
              className="skeleton"
              style={{
                height: "0.75rem",
                width: "60%",
                margin: "0 auto 0.75rem",
              }}
            />
            <div
              className="skeleton"
              style={{ height: "1.25rem", width: "40%", margin: "0 auto" }}
            />
          </div>
        ))}
      </div>
    );
  }

  console.log(data);

  if (!data) return null;

  return (
    <div>
      <p className="section-label">5-Day Forecast</p>
      <div className="forecast-strip">
        {data?.Days?.map((day) => (
          <div key={day.Date} className="forecast-day-card">
            <p className="forecast-date">{formatDate(day.Date)}</p>
            <p
              style={{
                fontSize: "0.75rem",
                color: "var(--text-muted)",
                margin: "0.375rem 0",
                textTransform: "capitalize",
              }}
            >
              {day.Condition}
            </p>
            <div className="forecast-temps">
              <span className="forecast-high">{Math.round(day.High)}°</span>
              <span className="forecast-low">{Math.round(day.Low)}°</span>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
