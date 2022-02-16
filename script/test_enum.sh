#!/bin/bash

grpcurl -plaintext -d @ 127.0.0.1:3000 enum.EnumService/Run <<EOM
{
"id": "user1",
"domain": "google.com",
"config": {
    "mode": 1,
    "timeout": 5
}
}
EOM
