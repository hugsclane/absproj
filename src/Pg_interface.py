from psycopg import Connection
import json
import requests
from rest import get_req
from models import Metadata
from src.crud import CRUD



class Pg_service:

    def __init__(self,host, port, dbname):
        self.conn_info = make_conninfo(host = host, port = port, dbname = dbname)
            # TODO: test me

    def create_connection():
        with Connection.connect(conninfo = self.conninfo , options="-c searchpath=demo") as conn:
            crud = CRUD(conn)
            return crud

# abandoned this until validation set up since allows sql injection.
   #def insert_metadata_stmt(table_name,**kwargs):
    #    #TODO validate kwargs with table_name
    #    keys = tuple(kwargs.keys())
    #    stmt = f"""
    #    INSERT INTO ({table_name}) ({keys})
    #    """


def main():
    with open("../metadata/abs_dataflow.json", 'r') as file:
        data = json.loads(file.read())
        file.close()
    print(data["data"])
    pg_serve = Pg_service()
    crud = pg_serve.create_connection(dbname = "absData",user = "admin",
                                      password = "admin", port = "5432")
    crud.insert_metadata(data["data"])
if __name__ == "__main__":
    main()



