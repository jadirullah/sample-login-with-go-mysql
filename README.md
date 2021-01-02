# sample-login-with-go-mysql

ini adalah simple aplikasi login dan register dengan Go serta database Mysql,

berikut petunjuk penggunaanya :) 


##Instalasi

- Buat database ber-nama **go_db** dan sebuah tabel di database mysql dengan nama tabel **users**

```
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `first_name` varchar(200) NOT NULL,
  `last_name` varchar(200) NOT NULL,
  `password` varchar(120) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

```

- lanjut ketikan perintah ini di terminal atau di git command

 ```git clone https://github.com/jadirullah/sample-login-with-go-mysql.git```

- masuk kedalam direktory sample-login-with-go-mysql atau direktory project
  ```cd sample-login-with-go-mysql```
  
- install package mysql
  ```go get -u github.com/go-sql-driver/mysql```
  
- install package golang.org/x/crypto/bcrypt
  ```go get -u golang.org/x/crypto/bcrypt```

- terakhir, jalankan perintah di bawah ini untuk running program

 ```go run main.go ```



##Success :D

##Semoga Bermanfaat
