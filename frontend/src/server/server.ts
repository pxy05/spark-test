import type { TodoItem } from "../lib/types";

const API_ADDRESS = "http://localhost:8080";
const API_TODO_ADDRESS = "http://localhost:8080/";

async function fetchTodos(): Promise<TodoItem[] | undefined> {
  try {
    const response = await fetch("http://localhost:8080/");
    if (response.status !== 200) {
      console.error("Error fetching data. Response status not 200");
      return undefined;
    }

    return (await response.json()) as TodoItem[]; //changed return type
  } catch (e) {
    console.error("Could not connect to server. Ensure it is running.", e);
  }
}

async function postTodos(todoItem: TodoItem): Promise<TodoItem | undefined> {
  if (todoItem === undefined) {
    console.error("todoItem is undefined.");
    return undefined;
  }

  if (todoItem.title === "") {
    console.error("todoItem must have a valid title.");
    return undefined;
  }

  if (todoItem.description === "") {
    console.error("todoItem must have valid description");
    return undefined;
  }

  try {
    const response = await fetch(API_TODO_ADDRESS, {
      method: "POST",
      body: JSON.stringify(todoItem),
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      console.error("Error with POST request for: ", todoItem, response.status);
      return undefined;
    }

    return (await response.json()) as TodoItem;
  } catch (e) {
    console.error("Error: Posting todoItem", todoItem, e);
    return undefined;
  }
}

export { fetchTodos, postTodos };
