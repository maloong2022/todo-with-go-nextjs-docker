"use client";

import Link from "next/link";
import { useRouter } from "next/navigation";
import React, { FormEvent } from "react";

const creteTodo = async (formData: FormData) => {
  const res = await fetch("http://localhost:8000/todos", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      Title: formData.get("title"),
      Content: formData.get("content"),
    }),
  });
  if (!res.ok) {
    throw new Error("Failed to create todo");
  }
  return res.json();
};

const CreatePost = () => {
  const router = useRouter();
  const handleCreateTodo = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    await creteTodo(formData);
    router.push("/");
  };
  return (
    <div>
      <form
        onSubmit={async (e) => {
          await handleCreateTodo(e);
        }}
        className="flex items-center flex-col  justify-center p-10"
      >
        <label className="form-control w-full max-w-xs">
          <div className="label">
            <span className="label-text">What is your todo title?</span>
          </div>
          <input
            type="text"
            name="title"
            placeholder="Just do IT"
            className="input input-bordered w-full max-w-xs"
          />
        </label>
        <label className="form-control w-full max-w-xs mt-5">
          <div className="label">
            <span className="label-text">What is your todo content?</span>
          </div>
          <textarea
            name="content"
            className="textarea textarea-bordered"
            placeholder="Software Development is a amazing life style"
          ></textarea>
        </label>
        <div className="mt-5 w-full max-w-xs flex justify-end">
          <Link href="/">
            <button className="btn bg-gray-600 mr-2 text-white">Back</button>
          </Link>

          <button type="submit" className="btn btn-success text-white">
            Save
          </button>
        </div>
      </form>
    </div>
  );
};

export default CreatePost;
