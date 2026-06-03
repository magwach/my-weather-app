export interface CurrentWeather {
  City: string;
  Temperature: number;
  Condition: string;
  Humidity: number;
  WindSpeed: number;
}

export interface ForecastDay {
  Date: string;
  High: number;
  Low: number;
  Condition: string;
}

export interface Forecast {
  City: string;
  Days: ForecastDay[];
}

export interface FavoriteWithWeather {
  city: string;
  weather: CurrentWeather;
}

export interface AuthResponse {
  token: string;
  name: string;
}

export interface AuthUser {
  name: string;
}