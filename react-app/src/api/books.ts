import axios from "axios";
import { Book } from "../types/book";

const API_URL = process.env.REACT_APP_API_URL || "http://localhost:8080";

export const getBooks = async (): Promise<Book[]> => {
  const response = await axios.get<Book[]>(`${API_URL}/books`);
  return response.data;
};
