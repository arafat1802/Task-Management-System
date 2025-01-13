package models

var CreateUser = `INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`

var GetUser = `SELECT * FROM users`