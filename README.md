# Credit card transaction - using Go (golang)
<p>Credit card transaction is a basic full stack web application that using this app user can buy widgets and monthly subscription using their credit card.</p>

### Tech Stack <a name="tech-stack"></a>

I used Go (golang), MariaDB, Bootstrap, Html, Javascript and css to build this project.
  <summary>Full Stack</summary>
  <ul>
    <li>Go</li>
    <li>MariaDB</li>
    <li>Bootstrap</li>
    <li>JAVASCRIPT</li>
    <li>Html</li>
    <li>CSS</li>
  </ul>

 <summary>Major Dependencies</summary>
  <ul>
  <li><a href="https://github.com/stripe/stripe-go/v72">WebSocket</a> To accept credit card transaction</li>
    <li><a href="https://github.com/go-chi/chi">Chi router</a></li>
    <li><a href="https://github.com/alexedwards/scs/v2">Alex edwards SCS </a> Session Manager</li>
    <li><a href="https://github.com/go-sql-driver/mysql">MySQL Driver</a> Database Driver</li>
    <li><a href="https://github.com/xhit/go-simple-mail/v2">simple mail</a> To create a simple mail server</li>
    <li><a href="https://github.com/gorilla/websocket">WebSocket</a> To Connect this app with websocket</li>
    
  </ul>

## Demo
![Capture](https://user-images.githubusercontent.com/35267447/227715745-dbc6cc12-7042-4fdb-a893-fade76add798.PNG)



## 💻 Getting Started
- To get star with this package first of all you have to clone the project ⬇️
``` bash
https://github.com/raihan2bd/go-credit-transact.git
```
- Then Make sure you have install [Go (golang)](https://go.dev/dl/) version 1.8.0 or latest stable version.
- Then make sure you have install [MariaDB](https://mariadb.org/) on your local mechine if you want to use this project as localy.
- To install all the Go packages navigate the folder address on your terminal and enter the below command ⬇️
``` bash
go mod tidy
```
- After downloading the packages you should rename example.database.yml file name to database.yml file and edit database credentials to your own database information and you have to create a `Makefile` inside the root as like example make file and provide your stripe key and secret as well.
![Capture5](https://user-images.githubusercontent.com/35267447/223344475-c64994c5-8a73-44d7-a571-5d3247c8db74.PNG)
- To setup database tables and columns by onClick install [soda cli database migration tool](https://gobuffalo.io/documentation/database/soda/) Then run below command ⬇️
```sh
soda migrate
```


# Usages
> *Note: Before enter the below command make sure you are in the right directory.*

- After finishing the avove instructions you can see the project in your local mechine by entering the below command ⬇️
```bash
make start
```
or 
```sh
go run cmd/web/* cmd/api/* cmd/micro/*
```

- Then you can see this project live on your browser by this link http://localhost:8080 or your given the port nuber you set for the project.


## 👥 Author

👤 **Abu Raihan**

- GitHub: [@githubhandle](https://github.com/raihan2bd)
- Twitter: [@twitterhandle](https://twitter.com/raihan2bd)
- LinkedIn: [LinkedIn](https://linkedin.com/in/raihan2bd)

## 🙏 Acknowledgments <a name="acknowledgements"></a>

I would like to thatnks [Trevor Sawler](https://github.com/tsawler) Who help me a lot learn go with this project. 

## ⭐️ Show your support <a name="support"></a>

> Thanks for visiting my repository. Give a ⭐️ if you like this project!

## 📝 License <a name="license"></a>

This project is [MIT](./LICENSE) licensed.

## Contribution
*Your suggestions will be more than appreciated. If you want to suggest anything for this project feel free to do that. :slightly_smiling_face:*
