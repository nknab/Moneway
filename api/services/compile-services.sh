#!/usr/bin/env bash
/*
 * File: compile-services.sh
 * Project: Moneway Go Developper Intern Challenge
 * File Created: Friday, 15th March 2019 6:53:31 PM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.1
 * Brief: 
 * -----
 * Last Modified: Friday, 15th March 2019 6:53:38 PM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

protoc -I . balance/pb/balance.proto --go_out=plugins=grpc:.
protoc -I . transaction/pb/transaction.proto --go_out=plugins=grpc:.