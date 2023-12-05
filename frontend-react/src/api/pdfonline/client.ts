import { AxiosRequestConfig } from "axios";

import { client, pickData } from "../client";
import { HTTPError, parseErrorMessage } from "../error";

export function PDFOnlineClient<T>(config: AxiosRequestConfig) {
  return client
    .request<T>(config)
    .then(pickData)
    .catch((error) => {
      console.log(error?.request?.url)
      throw new HTTPError(parseErrorMessage(error), error?.response?.status);
    });
}
