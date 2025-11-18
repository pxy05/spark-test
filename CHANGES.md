# Assumptions

## General

- Assumed that both title, and description fields must have some sort of value in them as the openAPI description specified them as required.

## Backend

- I assumed that the ToDoListHandler was implying the naming convention for handler functions -> ToDoHandler
- Therefore I routed them through a middleware function in order to have control over the responses to http methods that dont exist.
- Assumed there was a chance of simultaneous reads/writes therefore implemented read/write mutex.
- additionalProperties was set to false. I would have implemented some sort of UUID counter, and status property e.g. unfinished, finished, in progress...

## Frontend

- One thing to note is that "event?.preventDefault()" was used to prevent the page refreshing on every todo submission. However, event seems to be deprecated. If I had more time I would have found an undeprecated approach.
- If I had more time I would have added a bit more styling especially to the ugly scrollbar on the todo-list.

# General

## Backend

- Backend took around 40mins to write but adhering strictly to the openAPI spec and reading on openAPI/Golang "net/http" took an additional 30mins of triple checking and reading.
- Tests took around 10 minutes to write and passed first time
- Backend -> frontend debugging took around 15mins especially with adding "Content-Type": "application" to allowed headers and "Access-Control-Allow-Methods", "GET, POST, OPTIONS".
- API endpoint functionality tested with postman
- Race conditions mitigated with mutex locks. Mutex probably wasnt needed in this case as I assume its just for one person use however it is still good practice to implement.
- I would have like to implement some sort of logging in general via @ https://pkg.go.dev/github.com/mongodb/grip/slogger

## Frontend

- Svelte documentation for forms was easy enough although got mixed up with svelte/kit and svelte which cost some time.
- Spent a small bit of time with the colour palette and shadows @ https://html-css-js.com/css/generator/box-shadow/ to make everything look more visually appealing.
- Reordered the components so that they are inline and not stacked on eachother + max-size for todo-list to prevent page elongation.
- Text wrapping for long descriptions on todos.
