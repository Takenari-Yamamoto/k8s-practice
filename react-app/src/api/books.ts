import axios from "axios";
import { Book } from "../types/book";

// Kubernetes環境では、サービス名で接続
const API_URL = process.env.REACT_APP_API_URL || "http://golang-app:8080";

export const getBooks = async (): Promise<Book[]> => {
  const response = await axios.get<Book[]>(`${API_URL}/books`);
  return response.data;
};
