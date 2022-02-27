#!/bin/bash

grpcurl -plaintext -d @ 127.0.0.1:3000 viz.VizService/GetGraphistry <<EOM
{
"id": "user1",
"domain": "discord.com"
}
EOM
