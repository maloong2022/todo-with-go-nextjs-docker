"use client";

import React, { useEffect, useState } from "react";
import Link from "next/link";

const getData = async () => {
  const res = await fetch("http://localhost:8000/todos", {
    method: "GET",
    headers: {
      "Content-Type": "application.json",
    },
  });

  if (!res.ok) {
    throw new Error("Failed to get todo data");
  }

  return res.json();
};

const ListAllTodo: React.FC = () => {
  const [todos, setTodos] = useState([]);
  useEffect(() => {
    const getTodos = async () => {
      const todosInDB = await getData();
      console.log(todosInDB);
      setTodos(todosInDB);
    };
    getTodos();
  }, []);
  return (
    <div className="flex justify-center items-center mt-3 overflow-x-auto">
      <table className="table">
        {/* head */}
        <thead>
          <tr>
            <th></th>
            <th>Title</th>
            <th>Content</th>
            <th>CreatedAt</th>
          </tr>
        </thead>

        <tbody>
          {todos.map((todo: any) => (
            <tr key={todo.ID}>
              <td>{todo.ID}</td>
              <td>
                <Link
                  href={`/details/${todo.ID}`}
                  className="cursor-pointer font-bold text-red-500"
                >
                  {todo.Title}
                </Link>
              </td>
              <td>{todo.Content}</td>
              <td>{todo.Createdat}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default ListAllTodo;
