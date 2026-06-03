"use client";

import { Droplets, Wind } from "lucide-react";
import { CurrentWeather } from "@/types";
import { Star } from "lucide-react";
import { useAddFavorite, useFavorites } from "@/hooks/useFavorite";

interface Props {
  data: CurrentWeather | undefined;
  loading: boolean;
  error: any;
  city: string;
}

export default function WeatherHero({ data, loading, error, city }: Props) {
  const { mutate: addFavorite, isPending } = useAddFavorite();
  const { data: favorites } = useFavorites();

  const isFavorited = favorites?.some(
    (f) => f.city?.toLowerCase() === (data?.City ?? city).toLowerCase()
  );

  if (loading) {
    return (
      <div className="weather-hero">
        <div className="skeleton" style={{ height: "2rem", width: "40%", marginBottom: "0.75rem" }} />
        <div className="skeleton" style={{ height: "4rem", width: "25%", marginBottom: "0.5rem" }} />
        <div className="skeleton" style={{ height: "1rem", width: "30%" }} />
      </div>
    );
  }

  if (error || !data) {
    return (
      <div className="weather-hero">
        <p style={{ color: "var(--danger)", fontSize: "0.9rem" }}>
          {error?.response?.data?.error ?? "City not found. Please try another search."}
        </p>
      </div>
    );
  }

  return (
    <div className="weather-hero">
      <div style={{ display: "flex", justifyContent: "space-between", alignItems: "flex-start" }}>
        <div>
          <p className="weather-city">{data.City}</p>
          <p className="weather-condition">{data.Condition}</p>
        </div>

        {!isFavorited && (
          <button
            className="btn btn-ghost"
            style={{ gap: "0.4rem", fontSize: "0.8rem" }}
            onClick={() => addFavorite(data.City.toLowerCase())}
            disabled={isPending}
          >
            <Star size={15} />
            Save
          </button>
        )}

        {isFavorited && (
          <div style={{ display: "flex", alignItems: "center", gap: "0.4rem", fontSize: "0.8rem", color: "var(--accent)" }}>
            <Star size={15} fill="currentColor" />
            Saved
          </div>
        )}
      </div>

      <p className="weather-temp" style={{ margin: "1.25rem 0 1.5rem" }}>
        {Math.round(data.Temperature)}°
      </p>

      <div style={{ display: "flex", gap: "2rem" }}>
        <div className="weather-stat">
          <span className="weather-stat-label">Humidity</span>
          <span className="weather-stat-value" style={{ display: "flex", alignItems: "center", gap: "0.35rem" }}>
            <Droplets size={15} color="var(--accent)" />
            {data.Humidity}%
          </span>
        </div>
        <div className="weather-stat">
          <span className="weather-stat-label">Wind</span>
          <span className="weather-stat-value" style={{ display: "flex", alignItems: "center", gap: "0.35rem" }}>
            <Wind size={15} color="var(--accent)" />
            {data.WindSpeed} m/s
          </span>
        </div>
      </div>
    </div>
  );
}