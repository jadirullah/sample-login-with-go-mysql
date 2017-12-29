# sample-login-with-go-mysql

ini adalah simple aplikasi login dan register dengan Go serta database Mysql,

berikut petunjuk penggunaanya :) 


##Instalasi

- Buat database ber-nama **go_db** dan sebuah tabel di database mysql dengan nama tabel **users**

```
CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(50) DEFAULT NULL,
  `first_name` varchar(200) NOT NULL,
  `last_name` varchar(200) NOT NULL,
  `password` varchar(120) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

```

- lanjut ketikan perintah ini di terminal atau di git command

 ```git clone https://github.com/jadirullah/sample-login-with-go-mysql.git```

- terakhir, masuk ke dalam folder project go jalankan perintah 

 ```go run main.go ```



##Success :D

##Semoga Bermanfaat