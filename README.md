# MonewayV1.0
This is an intern assessment for the position of a go developper at Moneway.

This is a simulation of a transaction and update of a balance. There are two services, "balance" and "transaction". The "balance" service is responsible for handling the balance within the database, and the "transaction" service is responsible for updating the balance (by contacting the balance service) and storing the transaction data.

These are the Features:
1. list the balance
2. Credit and debit an Account
3. Create transactions
4. Update the balance 

The goal of the exercise is to apply (at most) the good practices of microservices ([microservices.io](https:microservices.io) and [gRPC](grpc.io)), the mastery of the Go language, and the mastery of database operations.

## Database
This makes use of MySql Database

In the **config/config.toml** file you will find the configuration file for the Database. Mkae changes to suit your database settings. 

`server = "127.0.0.1"
port = "3306"
database = "moneway"
user = "your username"
password = "your password"
`

The database schema could be found in the **moneway.sql** file
