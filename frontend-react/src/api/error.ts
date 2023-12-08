import { AxiosError } from "axios";

export class HTTPError extends Error {
  public status: number;

  constructor(message?: string, status?: number) {
    if (status === 404) {
      super("resource not found");
    } else {
      super(message || "internal server error");
    }
    this.status = status || 500;
  }
}

interface APIError {
  error?: string;
  errors?: { detail?: string }[];
}

export const parseErrorMessage = (error: AxiosError<APIError>) =>
  error?.response?.data?.error ||
  error?.response?.statusText;
