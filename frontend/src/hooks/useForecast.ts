import { useQuery } from "@tanstack/react-query";
import api from "@/lib/axios";
import { Forecast } from "@/types";

export const useForecast = (city: string) => {
  return useQuery({
    queryKey: ["forecast", city],
    queryFn: async () => {
      const { data } = await api.get<Forecast>(`/forecast/${city}`);
      return data;
    },
    enabled: !!city,
    staleTime: 1000 * 60 * 30,
    retry: false,
  });
};

export const useForecastByCoords = (lat: number | null, lon: number | null) => {
  return useQuery({
    queryKey: ["forecast", "coords", lat, lon],
    queryFn: async () => {
      const { data } = await api.get<Forecast>(`/forecast/coords`, {
        params: { lat, lon },
      });
      return data;
    },
    enabled: lat !== null && lon !== null,
    staleTime: 1000 * 60 * 30,
    retry: false,
  });
};