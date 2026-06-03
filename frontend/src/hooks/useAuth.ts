"use client";

import { useMutation } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
import { toast } from "sonner";
import api from "@/lib/axios";
import { setToken, setUser, logout } from "@/lib/auth";
import { AuthResponse } from "@/types";

interface RegisterPayload {
  name: string;
  email: string;
  password: string;
}

interface LoginPayload {
  email: string;
  password: string;
}

export const useRegister = () => {
  const router = useRouter();

  return useMutation({
    mutationFn: async (payload: RegisterPayload) => {
      const { data } = await api.post<AuthResponse>("/auth/register", payload);
      return data;
    },
    onSuccess: (data) => {
      setToken(data.token);
      setUser({ name: data.name });
      toast.success(`Welcome, ${data.name}!`);
      router.push("/dashboard");
    },
    onError: (error: any) => {
      const message = error.response?.data?.error || "Registration failed";
      toast.error(message);
    },
  });
};

export const useLogin = () => {
  const router = useRouter();

  return useMutation({
    mutationFn: async (payload: LoginPayload) => {
      const { data } = await api.post<AuthResponse>("/auth/login", payload);
      return data;
    },
    onSuccess: (data) => {
      setToken(data.token);
      setUser({ name: data.name });
      toast.success(`Welcome back, ${data.name}!`);
      router.push("/dashboard");
    },
    onError: (error: any) => {
      const message = error.response?.data?.error || "Login failed";
      toast.error(message);
    },
  });
};

export const useLogout = () => {
  const router = useRouter();

  return () => {
    logout();
    toast.success("Logged out successfully");
    router.push("/login");
  };
};