# Bank Server

Backend Server for a bank Infrastructure (Go and Postgres)

## Things to do

- [x] Design Database Schema
      ![Table Schema](bank-server.png)
- [x] Create a postgres instance using docker and docker volume
      using `docker-compose.yaml` file
- [x] Connect to a Postgres instance
- [x] Create Tables In DB using sql file
      Future Improvements:
  - [ ] use SQLC to make queries
- [ ] Make CRUD API for accounts table
- [ ] Write Unit test for the API's
- [ ] Dockerize (create a docker image) for the go app and db
- [ ] Create CI using Github Actions on master branch
- [ ] Handle Transactions
