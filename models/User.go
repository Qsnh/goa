package models

struct User {
    id int
    nickname string
    email string
    password string
    is_lock int8
    created_at int
    updated_at int
}