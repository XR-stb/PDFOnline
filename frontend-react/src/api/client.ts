import axios, { AxiosResponse } from "axios";

export const pickData = <T>(response: AxiosResponse<T>) => response.data;

const API_URL = "http://localhost:8080";

export const client = axios.create({
  baseURL: `${API_URL}/v1`,
  headers: {
    Accept: "application/json",
  },
});
