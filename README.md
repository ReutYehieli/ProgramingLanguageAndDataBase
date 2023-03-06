# ProgramingLanguageAndDataBase
Operating Instructions:
Need to download:
1. Golang, mySQL.
2. install Gin :  $go get github.com/gin-gonic/gin 
3. Install gorm: $go get github.com/jinzhu/gorm 
4. For mySQL install :  $go get github.com/go-sql-driver/mysql

Setting mySQL:
To register DB localy: need to update in filer config/DataBase.go , the relevant details : DBName, User, Password and Host, Port according to your database configuration.

Running the project:
running $ go run main.go

The server is immdentily connect and ready for getting message requset.
From here you can send your message requset :
https://mbarsinai.com/files/bgu/2022a/miniproj/swagger/#/ 


Explantion on project:
Image of system Architecture: 
![image](https://user-images.githubusercontent.com/63544809/223094249-fcd387b0-77f1-4ba6-9ebb-fd45d0e518be.png)

Working with frameworks:
- As a part of the project, I chose to use 2 different libary.
  The first one is gin, that give us support to create web server with using API calling.
  The second is gorm, that give us support with implement orm and working with DB.
  
  -Different Abstraction:
  -Working with ORM. Work with same objects in table columns in DB and implemeting SQL requset as part of GO language.
  -Rougting the API calls to a relevant function that hangles each call separately.
  
  Handle error:
  - DB throw exception if 2 different users have the same email.
  - Uniq ID : Using uuid.

