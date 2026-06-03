"use client";

import { useFavorites, useRemoveFavorite } from "@/hooks/useFavorite";
import { Trash2, Star } from "lucide-react";

interface Props {
  onSelectCity: (city: string) => void;
}

export default function FavoritesList({ onSelectCity }: Props) {
  const { data: favorites, isLoading } = useFavorites();
  const { mutate: removeFavorite } = useRemoveFavorite();

  console.log(favorites)

  if (isLoading) {
    return (
      <div>
        <p className="section-label">Favourites</p>
        <div
          style={{ display: "flex", flexDirection: "column", gap: "0.625rem" }}
        >
          {[1, 2].map((i) => (
            <div
              key={i}
              className="skeleton"
              style={{ height: "60px", borderRadius: "var(--radius-md)" }}
            />
          ))}
        </div>
      </div>
    );
  }

  if (!favorites || favorites.length === 0) {
    return (
      <div>
        <p className="section-label">Favourites</p>
        <div
          className="card-sm"
          style={{ textAlign: "center", padding: "1.5rem" }}
        >
          <Star
            size={20}
            color="var(--text-muted)"
            style={{ margin: "0 auto 0.5rem" }}
          />
          <p style={{ fontSize: "0.85rem", color: "var(--text-muted)" }}>
            No favourites yet. Search a city and save it.
          </p>
        </div>
      </div>
    );
  }

  return (
    <div>
      <p className="section-label">Favourites ({favorites.length}/3)</p>
      <div
        style={{ display: "flex", flexDirection: "column", gap: "0.625rem" }}
      >
        {favorites.map((fav) => (
          <div key={fav.city} className="favorite-item">
            <button
              onClick={() => onSelectCity(fav.city)}
              style={{
                background: "none",
                border: "none",
                cursor: "pointer",
                textAlign: "left",
                flex: 1,
              }}
            >
              <p
                style={{
                  fontWeight: 600,
                  fontSize: "0.95rem",
                  color: "var(--text-primary)",
                  textTransform: "capitalize",
                }}
              >
                {fav.city}
              </p>
              {fav.weather && (
                <p
                  style={{
                    fontSize: "0.8rem",
                    color: "var(--text-muted)",
                    marginTop: "0.15rem",
                    textTransform: "capitalize",
                  }}
                >
                  {Math.round(fav.weather.Temperature)}° ·{" "}
                  {fav.weather.Condition}
                </p>
              )}
            </button>

            <button
              className="btn-icon btn-danger"
              onClick={() => removeFavorite(fav.city)}
              title="Remove favourite"
            >
              <Trash2 size={15} />
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}
