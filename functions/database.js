/**
 * M3O Database
 * -------------------------
 *
 * This is a snippet of how you may want to consume the db api.
 */

import { call } from '..'

// ------------------------------
// Example 1 -> Individual Functions
// ------------------------------

export function count(table) {
  return call('db', 'Count', {
    table
  })
}

export function create(table, record) {
  return call('db', 'Create', {
    table,
    record
  })
}

export function deleteItem(table, id) {
  return call('db', 'Delete', {
    table,
    id
  })
}

export function dropTable(table) {
  return call('db', 'DropTable', { table })
}

export function read(table, query) {
  return call('db', 'Read', {
    table,
    query
  })
}

export function rename(from, to) {
  return call('db', 'Rename', {
    from,
    to
  })
}

export function truncate(table) {
  return call('db', 'Truncate', { table })
}

export function update(table, record) {
  return call('db', 'Update', {
    table,
    record
  })
}

// ------------------------------
// Database Class
// ------------------------------

class Database {
  constructor(table) {
    this.table = table
  }

  count() {
    return call('db', 'Count', {
      table: this.table
    })
  }

  create(record) {
    return call('db', 'Create', {
      table: this.table,
      record
    })
  }

  deleteItem(id) {
    return call('db', 'Delete', {
      table: this.table,
      id
    })
  }

  dropTable() {
    return call('db', 'DropTable', { table: this.table })
  }

  read(query) {
    return call('db', 'Read', {
      table: this.table,
      query
    })
  }

  async rename(to) {
    await call('db', 'Rename', {
      from: this.table,
      to
    })

    this.table = to
  }

  truncate() {
    return call('db', 'Truncate', { table: this.table })
  }

  update(record) {
    return call('db', 'Update', {
      table: this.table,
      record
    })
  }
}

// Example usage
// todo.js
/*
import { Database } from './db'

const todos = new Database('todos')

todos.create({
  todo,
  checked: false
})
*/
