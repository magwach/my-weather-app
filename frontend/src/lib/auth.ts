import Cookies from "js-cookie";

const TOKEN_KEY = "token";
const USER_KEY = "user";

export interface AuthUser {
  name: string;
}

export const getToken = (): string | undefined => {
  return Cookies.get(TOKEN_KEY);
};

export const setToken = (token: string): void => {
  Cookies.set(TOKEN_KEY, token, {
    expires: 1, 
    sameSite: "strict",
    secure: process.env.NODE_ENV === "production",
  });
};

export const removeToken = (): void => {
  Cookies.remove(TOKEN_KEY);
};

export const getUser = (): AuthUser | null => {
  const user = Cookies.get(USER_KEY);
  if (!user) return null;
  try {
    return JSON.parse(user);
  } catch {
    return null;
  }
};

export const setUser = (user: AuthUser): void => {
  Cookies.set(USER_KEY, JSON.stringify(user), {
    expires: 1,
    sameSite: "strict",
    secure: process.env.NODE_ENV === "production",
  });
};

export const removeUser = (): void => {
  Cookies.remove(USER_KEY);
};

export const isAuthenticated = (): boolean => {
  return !!getToken();
};

export const logout = (): void => {
  removeToken();
  removeUser();
};