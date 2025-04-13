import React, { useEffect, useState } from "react";
import { Book } from "../types/book";
import { getBooks } from "../api/books";

const BookList: React.FC = () => {
  const [books, setBooks] = useState<Book[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        const data = await getBooks();
        setBooks(data);
        setError(null);
      } catch (err) {
        setError("本の取得に失敗しました。");
      } finally {
        setLoading(false);
      }
    };

    fetchBooks();
  }, []);

  if (loading) return <div>読み込み中...</div>;
  if (error) return <div style={{ color: "red" }}>{error}</div>;

  return (
    <div>
      <h1>本の一覧</h1>
      <div style={{ display: "grid", gap: "1rem", padding: "1rem" }}>
        {books.map((book) => (
          <div
            key={book.id}
            style={{
              border: "1px solid #ccc",
              padding: "1rem",
              borderRadius: "4px",
            }}
          >
            <h2>{book.title}</h2>
            <p>著者: {book.author}</p>
            <p>出版年: {book.year}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default BookList;
