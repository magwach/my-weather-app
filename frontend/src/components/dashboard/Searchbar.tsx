"use client";

import { useState } from "react";
import { Search, LocateFixed, Loader2 } from "lucide-react";

interface Props {
  value: string;
  onChange: (v: string) => void;
  onSearch: (city: string) => void;
  onUseLocation: () => void;
  geoLoading: boolean;
  geoError: string | null;
}

export default function SearchBar({
  value,
  onChange,
  onSearch,
  onUseLocation,
  geoLoading,
  geoError,
}: Props) {
  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter" && value.trim()) {
      onSearch(value.trim());
    }
  };

  return (
    <div style={{ display: "flex", flexDirection: "column", gap: "0.5rem" }}>
      <div style={{ display: "flex", gap: "0.75rem" }}>
        <div className="search-bar" style={{ flex: 1 }}>
          <Search size={16} color="var(--text-muted)" />
          <input
            placeholder="Search for cities..."
            value={value}
            onChange={(e) => onChange(e.target.value)}
            onKeyDown={handleKeyDown}
          />
          {value.trim() && (
            <button
              className="btn btn-primary"
              style={{ padding: "0.4rem 0.875rem", fontSize: "0.8rem" }}
              onClick={() => onSearch(value.trim())}
            >
              Search
            </button>
          )}
        </div>

        <button
          className="btn btn-ghost"
          onClick={onUseLocation}
          disabled={geoLoading}
          title="Use my location"
          style={{ gap: "0.5rem", whiteSpace: "nowrap" }}
        >
          {geoLoading ? (
            <Loader2 size={16} style={{ animation: "spin 1s linear infinite" }} />
          ) : (
            <LocateFixed size={16} />
          )}
          {geoLoading ? "Locating..." : "My location"}
        </button>
      </div>

      {geoError && (
        <p style={{ fontSize: "0.8rem", color: "var(--danger)", paddingLeft: "0.25rem" }}>
          {geoError}
        </p>
      )}
    </div>
  );
}