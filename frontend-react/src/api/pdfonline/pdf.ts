import { PDFOnlineClient } from "./client";
import { PdfType } from "../../types";

export const listPdfs = () =>
  PDFOnlineClient<{pdfs: PdfType[]}>({
    url: "pdfs",
    method: "get",
  }).then((data) => data.pdfs);