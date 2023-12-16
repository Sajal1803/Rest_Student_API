This a book rest api for a book store which stores the data of the books present in the bookstore and perform CRUD operations.

1. GET: we can get all the books available in the book store.
2. POST: Add new book to the database.
3. PUT: Update book details with given id.
4. DELETE: Delete a book with a given id, which is removed or sold from book store.


Sequence Diagram for this Book API:

![sequence Diagram]! (https://github.com/Sajal1803/Rest_Student_API/assets/76404926/8833bb40-a761-445f-a669-1a682e2c1451)



To initiate docker use these commands:
> docker pull mysql

> docker run --name student-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=student_db -p 3307:3306 -d mysql:8.0.30

> docker exec -it sample-mysql mysql -uroot -proot123 student_db -e "CREATE TABLE students (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), age INT, branch VARCHAR(255));"



We have used the default port i.e 8000

use 'http://localhost:8000/students' to access the GET and POST Query.

use 'http://localhost:8000/students/{id}' to access the PUT and DELETE Query.
