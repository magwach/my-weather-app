import { useState, useCallback } from "react";

interface GeolocationState {
  lat: number | null;
  lon: number | null;
  loading: boolean;
  error: string | null;
}

export const useGeolocation = () => {
  const [state, setState] = useState<GeolocationState>({
    lat: null,
    lon: null,
    loading: false,
    error: null,
  });

  const getLocation = useCallback(() => {
    if (!navigator.geolocation) {
      setState((prev) => ({
        ...prev,
        error: "Geolocation is not supported by your browser",
      }));
      return;
    }

    setState((prev) => ({ ...prev, loading: true, error: null }));

    navigator.geolocation.getCurrentPosition(
      (position) => {
        setState({
          lat: parseFloat(position.coords.latitude.toFixed(2)),
          lon: parseFloat(position.coords.longitude.toFixed(2)),
          loading: false,
          error: null,
        });
      },
      (error) => {
        const messages: Record<number, string> = {
          1: "Location access denied",
          2: "Location unavailable",
          3: "Location request timed out",
        };
        setState({
          lat: null,
          lon: null,
          loading: false,
          error: messages[error.code] || "Failed to get location",
        });
      },
      {
        timeout: 10000,
        maximumAge: 1000 * 60 * 5,
      },
    );
  }, []);

  const clearLocation = useCallback(() => {
    setState({ lat: null, lon: null, loading: false, error: null });
  }, []);

  return { ...state, getLocation, clearLocation };
};
