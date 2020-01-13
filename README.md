# todo_app
Create an todo app app with standard structure, Login with JWT and test as well.


### Todo app requirements

1. New user should register with (username, email_add, password).
2. User login using (email_add, password) and system return Token thorugh JWT.
3. User can get their relavent todo items with valid token. 
4. User can add more todo items with valid token.
5. User can update todo item with valid token.
6. User can remove todo item with valid token.
7. Make intergration tests of Login, Register, Add, Update, Delete methods.



### Dir structure

1. todo_app
   - auth
     - jwt.go
     - jwt_test.go
   - bin
     - todoapp
       - main.go
   - config
     - conf.go
   - controllers
      - auth.go
      - items.go
   - db
     - db_data.go   
   - routes
      - routes.go
   - store
      - todo.go
      - user.go
   - types
     - types.go
