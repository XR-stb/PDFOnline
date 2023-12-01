import { PDFOnlineClient } from "./client";

export interface Pdf {
  id: string;
  author: string;
  title: string;
  description: string;
  url: string;
  cover_url: string;
}

export const listPdfs = () =>
  PDFOnlineClient<{pdfs: Pdf[]}>({
    url: "pdfs",
    method: "get",
  }).then((data) => data.pdfs);