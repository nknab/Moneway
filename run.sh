 #
 # File: run.sh
 # Project: Moneway Go Developper Intern Challenge
 # File Created: Sunday, 17th March 2019 7:35:48 PM
 # Author: nknab
 # Email: kojo.anyinam-boateng@outlook.com
 # Version: 1.1
 # Brief: This script is to help the user run the various servers
 # -----
 # Last Modified: Sunday, 17th March 2019 7:39:28 PM
 # Modified By: nknab
 # -----
 # Copyright ©2019 nknab
 #

#!/usr/bin/env bash


 #
 # @brief Building and running the main server
 #
 # @return void
 #
function main-server(){
    go build
    ./Moneway
}

 #
 # @briefBuilding and running the balance-service server
 #
 # @return void
 #
function balance-server(){
    cd api/services/balance
    go build
    ./balance
}

 #
 # @briefBuilding and running the transaction-service server
 #
 # @return void
 #
function transaction-server(){
    cd api/services/transaction
    go build
    ./transaction
}