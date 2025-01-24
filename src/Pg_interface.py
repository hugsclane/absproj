from psycopg import Connection, conninfo
import json
import requests
from rest import get_req
from models import *
from crud import CRUD



class Pg_service:

    ##host omitted in conninfo for local testing
    def __init__(self,host, port, dbname,password, user):
        self.conn_info =conninfo.make_conninfo(host = host,port = port, dbname = dbname,
                                       user =user, password = password)
        # print(self.conn_info)
            # TODO: test me

    def create_connection(self):
        conn = Connection.connect(conninfo = self.conn_info)
        # , options="-c searchpath=db")
        crud = CRUD(conn)
        self.conn = conn
        return crud

    def close_connection(self):
        self.conn.close()

## TODO: this function is to be moved to the pw handler in the DB with an encryption step.


# abandoned this until validation set up since allows sql injection.
   #def insert_metadata_stmt(table_name,**kwargs):
    #    #TODO validate kwargs with table_name
    #    keys = tuple(kwargs.keys())
    #    stmt = f"""
    #    INSERT INTO ({table_name}) ({keys})
    #    """

#TODO: hugo run
# git filter-repo --invert-paths --path db_docs/pwdf
# at some point

def pwfromfile():
    with open("db_docs/pwdf.txt",'r', encoding = "UTF-8") as file:
        val = file.read()
        file.close()
        return val.strip("\n")

# currently only testing reference stubs from the metadata page.
def main():

    #pull data out of the metadata.json (simulating get request)
    with open("metadata/reference_stubs.json", 'r') as file:
        data = json.loads(file.read())
        file.close()
    #Process data into Metadata model class
    #TODO: NEXT JOB!!! go to explore.py and store all of the metadata.json data in the pydantic classes.
    #do whatever validation makes you feel warm inside and then spit it out to the just below here
    # loop over the number of unique fields (~1026) and then add them simultaneously to the DB.
    # after that we write tests for each function in preperation for building the flask api.

    pw = pwfromfile()
    # print(pw)
    # print(data["data"])
    pg_serve = Pg_service(host="localhost",dbname = "postgres",user = "hcl", password = pw, port = "5432")
    crud = pg_serve.create_connection()


    ## TODO: belongs in its own function, also handle i.pop("unwanted_fields")
    model_list = []
    unwanted_fields = ["names","descriptions","annotations","structure"]
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

        id = crud.insert_metadata(column)

        try:
            model_list.append(crud.get_metadata(id))
        except Exception as e:
            print(i)
    for m in model_list:
        print(m.model_dump())
    pg_serve.close_connection()
if __name__ == "__main__":
    main()



