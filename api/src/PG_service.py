from psycopg import Connection, conninfo
import json
import requests
from rest import get_req
from models import *
from crud import CRUD
from utils.local import *


class PG_service:

    ##host omitted in conninfo for local testing
    def __init__(self,host, port, dbname,password, user):
        self.conn_info =conninfo.make_conninfo(
                host = host,port = port, dbname = dbname,user =user, password = password)
        self.conn = create_connection()
        self.crud = create_crud()

    def create_connection(self):
        conn = Connection.connect(conninfo = self.conn_info)
        # don't autocommit the transactions or it will bugger up the blob.:w

        return conn
        # , options="-c searchpath=db")

    def create_crud(self):
        crud = CRUD(conn)
        return crud

    def close_connection(self):
        self.conn.close()



def main():
    pw = pwfromfile("db_docs/pwdf.txt")

    pg_serve = Pg_service(host="localhost",dbname = "postgres",user = "hcl", password = pw, port = "5432")

    conn = pg_serve.create_connection()


    crud = pg_serve.create_crud()

    model_list = []
    unwanted_fields = ["names","descriptions","annotations","structure"]
    data = of_json("metadata/reference_stubs.json")
    for i in data["data"]["dataflows"]:
        for j in unwanted_fields:
            try:
                i.pop(j)
            except:
                pass
            pass
        i["dataflowId"] = i.pop("id")
        i["agencyId"] = i.pop("agencyID")
        # print(i)
        # break
        column = MetadataInput(**i)
        # print(len(model_list))
        # print(column)
        # print(crud.insert_metadata(column))

        try:
            id = crud.insert_metadata(column)
            model_list.append(crud.get_metadata(id))
        except Exception as e:
            pass
    for m in model_list:
        print(m.model_dump())
    pg_serve.close_connection()

if __name__ == "__main__":
    main()



