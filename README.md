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

Moving forward as we progress we will be creating tasks and keeping a track on what we accomplished so far, where we go stuck, what bugs fucked up our night and where we outperformed.

P.S. We will not be using any framework, we will use mux from go.

## Day 1: (Sept 2, 2026)

### Tasks for Day 1:

1. Setup the folder structure. ✅
2. Create the server ✅
   - We faced a blocker here, we were not able to fetch the PORT from .env file. At that moment we realised we need godotenv package to load env variables from the .env file.

3. Setup graceful shutdown ✅
4. Try to use a go routine ✅

### Tasks for Day 2:

1. Create separate handlers ✅
2. Create logger middleware ✅
3. Create a postgres instant in docker ✅
4. Create migration folder to keep all the database migrations in one place ✅
5. Create a urls table to record url data ✅
6. Create config and database folder and write config and database files ✅
7. Bring the Pool connection in main.go ✅
8. Create models for CreateShortURL response and CreateShortURL request ✅

### Tasks for Day 3:

1. Create the infrastructure for our application ( handler , service , repository ) ✅
2. Create the repository structure ✅
3. Create the service structure
4. Create the handler structure
5. Create a basic POST /create api flow to create a shortURL.
