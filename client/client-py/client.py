"""The Python implementation of the GRPC enum.EnumServier.BasicEnumerate client."""

from __future__ import print_function

import logging
import sys

sys.path.insert(0, './gen/amass/enum')
sys.path.insert(0, './gen/amass')
sys.path.insert(0, './gen')

import grpc
import gen.amass.enum.enum_pb2 as enum
import gen.amass.enum.enum_pb2_grpc as pb


def run():
    with grpc.insecure_channel('127.0.0.1:3000') as channel:
        stub = pb.EnumServiceStub(channel)
        reponse = stub.BasicEnumerate(enum.EnumRequest(domains=["reg.chula.ac.th"]))
    print("EnumService client received: \n")
    print(reponse.results)

if __name__ == "__main__":
    logging.basicConfig()
    run()
