/**
 * M3O TODOs
 * -------------------------
 *
 * Make todo lists
 */

// todo.js

import { Database } from './db'

const todos = new Database('todos')

// Example of creating a todo
/*
  todos.create({
    todo,
    checked: false
  })
*/

import { create } from './db'

// createTodo is a short hand for creating a todo task
async function createTodo(todo) {
  const response = await create('todos', {
    todo,
    checked: false
  })

  console.log(response)
}

// Example usage
/*
  createTodo('Use M3O')
*/
