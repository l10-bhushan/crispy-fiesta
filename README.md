# Learning backend using go

## 1: Project structure

While building a go backend we need to have a proper folder structure. Starting with two main folders as **cmd** and **internal**.

## CMD - folder

This is where our server/main.go will live. We separate the main.go file
for readibility and testing

## Internal - folder

This folder will contain our entire application specs. This add a layer of encapsulation. There are no private or public keywords in Go. So moving code to different folders and files help us provide encapsulation and also we use dependency injection in Go. So, that every module has it's own responsibility and doesn't need to know everything.

For e.g:

1. Respository only knows how to communicate with postgres.
2. Service layer only handles the business logic
3. Handler layer only handles the Request and Response.
