Program API di mana terdapat dua grup endpoint, yaitu users dan photos
Users Endpoint
 1. POST : /users/register
   a. ID (primary key, required)
   b. Username (required)
   c. Email (unique & required)
   d. Password (required & minlength 6)
   e. Relasi dengan model Photo (Gunakan constraint cascade)
   f. Created At (timestamp)
   g. Updated At (timestamp)
 2. POST: /users/login http://localhost/users/login
   a. Using email & password (required)
 3. PUT : /users/:userId (Update User) (authMiddleware)
 4. DELETE : /users/:userId (Delete User) (authMiddleware)

Photos Endpoint
  1. POST : /photos (Upload photo)(authMiddleware)
   a. ID (primary key, required)
   b. Title
   c. Caption
   d. PhotoUrl
   e. UserID
   f. Relasi dengan model User
  2. GET : /photos (Get Photos)(authMiddleware)
  3. PUT : /photos/:photoId (Edit Photo) (authMiddleware)
  4. DELETE : /photos/:photoId (Remove Photo) (authMiddleware)

Progamming Language: Go
Database: MySQL
FrameWork: Gin Gonic
Tools: GORM, JWT, Go-Migrate, Gin-Contrib (Cors), Joho GoDotEnv

Tugas Akhir (Task 5) dari program Project Based Intership Fullstack Developer - Rakamin
