"use client";

import React, { useEffect, useState } from "react";
import DetailCard from "./DetailCard";

interface SinglePostProps {
  id: number;
}

const getSinglePostData = async (id: number) => {
  const res = await fetch(`http://localhost:8000/todos/${id}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    throw new Error("Failed to get the todo details");
  }
  return res.json();
};

const SinglePost: React.FC<SinglePostProps> = ({ id }: SinglePostProps) => {
  const [todo, setTodo] = useState();

  useEffect(() => {
    const getTodo = async () => {
      const todoInDB = await getSinglePostData(id);
      setTodo(todoInDB);
    };
    getTodo();
  }, [id]);
  return <div>{todo && <DetailCard todo={todo} />}</div>;
};

export default SinglePost;
