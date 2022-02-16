#!/bin/bash

grpcurl -plaintext -d @ 127.0.0.1:3000 db.DBService/Run <<EOM
{
"id": "user1",
"domain": "reg.chula.ac.th"
}
EOM