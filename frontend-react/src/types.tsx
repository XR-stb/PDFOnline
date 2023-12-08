export interface UserType {
  id: string;
  username: string;
  email: string;
  role: number;
  avatar: string;
}

export interface PdfType {
  id: string;
  author: string;
  title: string;
  description: string;
  url: string;
  cover_url: string;
}