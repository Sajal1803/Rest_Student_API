This a student rest api for a student database which stores the data of the students present in the college and perform CRUD operations.

1. GET: we can get all the Students available in the database.
2. POST: Add new students to the database.
3. PUT: Update student details with given id.
4. DELETE: Delete a student with a given id, which is removed or passed out from the college .


Sequence Diagram for this Student API:


![sequence  ](https://github.com/Sajal1803/Rest_Student_API/assets/76404926/181315ec-7809-448f-a45f-0c98f812520a)



To initiate docker use these commands:
> docker pull mysql

> docker run --name student-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=student_db -p 3307:3306 -d mysql:8.0.30

> docker exec -it sample-mysql mysql -uroot -proot123 student_db -e "CREATE TABLE students (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), age INT, branch VARCHAR(255));"



We have used the default port i.e 8000

use 'http://localhost:8000/students' to access the GET and POST Query.

use 'http://localhost:8000/students/{id}' to access the PUT and DELETE Query.
