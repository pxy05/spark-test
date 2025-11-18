<script lang="ts">
  import Todo from "./lib/Todo.svelte";
  import type { TodoItem } from "./lib/types";
  import { fetchTodos, postTodos } from "./server/server"

  let todos: TodoItem[] | undefined = $state([]);

  let todoTitle: string = $state("");
  let todoDescription: string = $state("");
  let todoError: string = $state("");

  const MISSING_TITLE_ERROR = "Title cannot be empty.";
  const MISSING_DESCRIPTION_ERROR = "Description cannot be empty.";
  const POST_ERROR = "Error creating TodoItem."

  function resetTodoFields() {
    todoTitle = "";
    todoDescription = "";
    todoError = "";
  }

  function updateTodoError(error: string) {
    console.log(error);
    todoError = error;
  }


async function handleSubmission(): Promise<void> {
  event?.preventDefault();
  const title = todoTitle.trim();
  const description = todoDescription.trim();

  if (title === "") {
    updateTodoError(MISSING_TITLE_ERROR)
    return undefined;
  }

  if (description === "") {
    updateTodoError(MISSING_DESCRIPTION_ERROR);
    return undefined;
  }

  const todoItem: TodoItem = {
    title: title,
    description: description,
  }

  let response: TodoItem | undefined;
  
  response = await postTodos(todoItem)


  if (response === undefined) {
    todoError = POST_ERROR;
    return;
  }

  console.log("successful post:", response)

  if (todos !== undefined) {
    todos.push(response);
  }
  
  resetTodoFields();

}

  // Initially fetch todos on page load
  $effect(() => {
    (async () => {
      todos = await fetchTodos();
    })();
  });
</script>

<main class="app">
  <header class="app-header">
    <h1>TODO</h1>
  </header>
<div class="todo-container">
    <div class="form-container default-card">
      <h2 class="todo-list-form-header">Add a Todo</h2>
      <form class="todo-list-form" onsubmit={handleSubmission}>
        <input bind:value={todoTitle} placeholder="Title" name="title" />
        <input bind:value={todoDescription} placeholder="Description" name="description" />
        <button type="submit" >Add Todo</button>
        <p class="form-error">{todoError}</p>
      </form>
    </div>
      <div class="todo-list">
      {#if todos === undefined}
        <Todo title="Could not connect to Server" description="" />
      {:else}
      {#each todos as todo}
        <Todo title={todo.title} description={todo.description} />
      {/each}

      {/if}

    </div>
  </div>
</main>