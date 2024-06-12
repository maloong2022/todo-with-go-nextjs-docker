"use client";

import Link from "next/link";
import { useRouter } from "next/navigation";

interface Todo {
  ID: number;
  Title: string;
  Content: string;
  Createdat: string;
}

interface DetailCardProps {
  todo: Todo;
}

const deleteTodo = async (id: number) => {
  const res = await fetch(`http://localhost:8000/todos/${id}`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    throw new Error("Failed to delete todo!");
  }
  return res.json();
};

const DetailCard: React.FC<DetailCardProps> = ({ todo }) => {
  const router = useRouter();

  const handleDelete = async () => {
    await deleteTodo(todo.ID);
    router.push("/");
  };
  return (
    <div className="card w-96 shadow-xl m-2">
      <div className="card-body">
        <h2 className="card-title uppercase">{todo.Title}</h2>
        <p>{todo.Content}</p>
        <p className="text-gray-400">{todo.Createdat}</p>
        <div className="card-actions justify-end mt-5">
          <Link href={`/edit/${todo.ID}`}>
            <button className="btn btn-primary">Edit</button>
          </Link>
          <button onClick={handleDelete} className="btn btn-error text-white">
            Delete
          </button>
        </div>
      </div>
    </div>
  );
};

export default DetailCard;
