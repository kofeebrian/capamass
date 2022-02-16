#!/bin/bash

grpcurl -plaintext -d @ 127.0.0.1:3000 enum.EnumService/Run <<EOM
{
"id": "user1",
"domain": "reg.chula.ac.th",
"config": {
    "mode": 0,
    "timeout": 5
}
}
EOM
