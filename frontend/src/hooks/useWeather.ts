import { useQuery } from "@tanstack/react-query";
import api from "@/lib/axios";
import { CurrentWeather } from "@/types";

export const useWeather = (city: string) => {
  return useQuery({
    queryKey: ["weather", city],
    queryFn: async () => {
      const { data } = await api.get<CurrentWeather>(`/weather/${city}`);
      console.log("hit")
      return data;
    },
    enabled: !!city,
    staleTime: 1000 * 60 * 10,
    retry: false,
  });
};

export const useWeatherByCoords = (lat: number | null, lon: number | null) => {
  return useQuery({
    queryKey: ["weather", "coords", lat, lon],
    queryFn: async () => {
      const { data } = await api.get<CurrentWeather>(`/weather/coords`, {
        params: { lat, lon },
      });
      return data;
    },
    enabled: lat !== null && lon !== null,
    staleTime: 1000 * 60 * 10,
    retry: false,
  });
};