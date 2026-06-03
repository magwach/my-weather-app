"use client";

import { useState } from "react";
import ProtectedRoute from "@/components/auth/ProtectedRoute";
import Sidebar from "@/components/dashboard/Sidebar";
import { useWeather, useWeatherByCoords } from "@/hooks/useWeather";
import { useForecast, useForecastByCoords } from "@/hooks/useForecast";
import { useGeolocation } from "@/hooks/useGeolocation";
import SearchBar from "@/components/dashboard/Searchbar";
import WeatherHero from "@/components/dashboard/WeatherHero";
import ForecastStrip from "@/components/dashboard/ForecastStrip";
import FavoritesList from "@/components/dashboard/FavouriteList";

export default function DashboardPage() {
  const [searchedCity, setSearchedCity] = useState("");
  const [activeCity, setActiveCity] = useState("");

  const {
    lat,
    lon,
    loading: geoLoading,
    error: geoError,
    getLocation,
    clearLocation,
  } = useGeolocation();

  const usingCoords = lat !== null && lon !== null && !activeCity;

  const {
    data: weatherByCity,
    isLoading: weatherCityLoading,
    error: weatherCityError,
  } = useWeather(usingCoords ? "" : activeCity);

  const { data: weatherByCoords, isLoading: weatherCoordsLoading } =
    useWeatherByCoords(usingCoords ? lat : null, usingCoords ? lon : null);

  const { data: forecastByCity, isLoading: forecastCityLoading } = useForecast(
    usingCoords ? "" : activeCity,
  );

  const { data: forecastByCoords, isLoading: forecastCoordsLoading } =
    useForecastByCoords(usingCoords ? lat : null, usingCoords ? lon : null);

  const weather = usingCoords ? weatherByCoords : weatherByCity;
  const forecast = usingCoords ? forecastByCoords : forecastByCity;
  const weatherLoading = usingCoords
    ? weatherCoordsLoading
    : weatherCityLoading;
  const forecastLoading = usingCoords
    ? forecastCoordsLoading
    : forecastCityLoading;

  const handleSearch = (city: string) => {
    clearLocation();
    setActiveCity(city);
    setSearchedCity(city);
  };

  const handleUseLocation = () => {
    setActiveCity("");
    setSearchedCity("");
    getLocation();
  };

  const handleFavoriteSelect = (city: string) => {
    clearLocation();
    setActiveCity(city);
    setSearchedCity(city);
  };

  return (
    <ProtectedRoute>
      <div className="dashboard-layout">
        <Sidebar />

        <main className="main-content">
          <SearchBar
            value={searchedCity}
            onChange={setSearchedCity}
            onSearch={handleSearch}
            onUseLocation={handleUseLocation}
            geoLoading={geoLoading}
            geoError={geoError}
          />

          <div
            style={{
              display: "flex",
              flexDirection: "column",
              gap: "1.25rem",
              marginTop: "1.5rem",
            }}
          >
            {(activeCity || usingCoords) && (
              <>
                <WeatherHero
                  data={weather}
                  loading={weatherLoading}
                  error={!usingCoords ? weatherCityError : null}
                  city={activeCity}
                />
                <ForecastStrip data={forecast} loading={forecastLoading} />
              </>
            )}

            {!activeCity && !usingCoords && !geoLoading && (
              <div
                className="card"
                style={{ textAlign: "center", padding: "3rem 1.5rem" }}
              >
                <p style={{ color: "var(--text-muted)", fontSize: "0.9rem" }}>
                  Search for a city or use your current location to get started.
                </p>
              </div>
            )}

            <FavoritesList onSelectCity={handleFavoriteSelect} />
          </div>
        </main>
      </div>
    </ProtectedRoute>
  );
}
