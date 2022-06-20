# Bank Server

Backend Server for a bank Infrastructure (Go and Postgres)

![docker image](https://github.com/prakharporwal/bank-server/actions/workflows/ecs-image-build.yml/badge.svg)

## RoadMap

- [x] Design Database Schema
      ![Table Schema](bank-server.png)
- [x] Create a postgres instance using docker and docker volume
      using `docker-compose.yaml` file
- [x] Connect to a Postgres instance
- [x] Create Tables In DB using sql file
      Future Improvements:
  - [ ] use SQLC to make queries
- [x] Make CRUD API for accounts table
- [x] Write Unit test for the API's
- [x] Dockerize (create a docker image) for the go app and db
      [See Dockerfile here](Dockerfile)
- [x] Create CI using Github Actions on master branch
  - [x] Push Docker Image to AWS ECR via actions
  - [x] Run go build and test
  - [x] Add Datree in CI for misconfig checking
- [x] Handle DB Transactions use the idea
```
BEGIN
if Succeed
    COMMIT
else
    ROLLBACK
```
- [x] Deploy on Kubernetes using AWS EKS and
    - [x] Install Kubecost for cost management
    - [x] install ArgoCD for GitOps

### TODO
- [ ] Setup Monitoring Using Prometheus
- [ ] Create Docker Network to let the image connect with the db
      `docker network create bank-network`
      `docker network connect bank-network postgres12`
- [ ] Read env variables from config file
- [ ] Write Unit Tests by using Mock to mock DB.


## How to Use ?

#### Pre-requisite :
- go1.17 or above installed
  - docker
  - make command
1. Install essential go modules
2. run `make postgres` for making a postgres db using DOCKER
3. Connect to the DB using `user: admin` and `password: password`
   and `db = default_db`.
4. run the [bank-server.sql](db/bank-server.sql) file to create db tables.
5. run `make server` the server will be up and running on localhost:8080

- POST /account for creating a new account

```json lines
{
    "owner_email":"sample@getmail.com",
    "currency":"USD"
}
```

- GET /account?owner_email=email@email.com for gettiing account details using email
- POST /transfer for transferring money to other bank account

```json lines
{
    "from_account_id": 2,
    "to_account_id": 1,
    "amount": 15
}
```