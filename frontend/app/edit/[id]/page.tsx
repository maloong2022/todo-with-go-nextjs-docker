"use client";

import EditPost from "@/components/EditPost";
import { Todo } from "@/type/Todo";
import React, { useEffect, useState } from "react";

const getData = async (id: number) => {
  const res = await fetch(`http://localhost:8000/todos/${id}`, {
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

const EditPage: React.FC = ({ params }: any) => {
  const { id } = params;
  const [todo, setTodo] = useState<Todo>();
  useEffect(() => {
    const getTodo = async () => {
      const todoInDB = await getData(id);
      setTodo(todoInDB);
    };
    getTodo();
  }, [id]);
  return <div>{todo && <EditPost todo={todo} />}</div>;
};

export default EditPage;
