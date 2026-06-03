import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";
import api from "@/lib/axios";
import { FavoriteWithWeather } from "@/types";

export const useFavorites = () => {
  return useQuery({
    queryKey: ["favorites"],
    queryFn: async () => {
      const { data } = await api.get<FavoriteWithWeather[]>("/favorites");
      return data;
    },
    staleTime: 1000 * 60 * 5,
  });
};

export const useAddFavorite = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (city: string) => {
      const { data } = await api.post("/favorites", { city });
      return data;
    },
    onSuccess: (_, city) => {
      queryClient.invalidateQueries({ queryKey: ["favorites"] });
      toast.success(`${city} added to favorites`);
    },
    onError: (error: any) => {
      const message = error.response?.data?.error || "Failed to add favorite";
      toast.error(message);
    },
  });
};

export const useRemoveFavorite = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (city: string) => {
      const { data } = await api.delete(`/favorites/${city}`);
      return data;
    },
    onSuccess: (_, city) => {
      queryClient.invalidateQueries({ queryKey: ["favorites"] });
      toast.success(`${city} removed from favorites`);
    },
    onError: (error: any) => {
      const message =
        error.response?.data?.error || "Failed to remove favorite";
      toast.error(message);
    },
  });
};