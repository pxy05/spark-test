<script lang="ts">
  import type { TodoItem } from "./types";

  const { title, description }: TodoItem = $props();
</script>

<div class="todo">
  <div class="todo-details">
    <p class="todo-title">{title}</p>
    <p class="todo-description">{description}</p>
  </div>
</div>

<style>
  .todo {
    display: flex;
    justify-content: space-between;
    align-items: center;
    color: rgb(92, 92, 92);

    padding: 15px;

    border-radius: 20px;
    border-width: 2px;
    border-style: solid;
    border-color: #9e9e9e65;
    background-color: #ffffff;


    max-width: 100%;

    margin: 10px 0px;
  }

  .todo-details {
    display: flex;
    flex-direction: column;
    align-items: start;

    max-width: 100%;
  }

  .todo-title {
    font-size: 28px;
    font-weight: bold;

    margin: 0px;
    margin-bottom: 5px;

    word-wrap: break-word;
    max-width: 100%;
  }

  .todo-description {
    display: flex;
    margin: 0px;

    font-size: 20px;

    text-align: left;
    word-break: break-all;
    max-width: 100%;
  }
</style>
