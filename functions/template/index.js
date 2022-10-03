const functions = require('@google-cloud/functions-framework')
const m3o = require('m3o').default
const { handler } = require('./handler/index.js')

global.m3o = m3o(process.env.NODE_ENV)

functions.http('main', handler)
