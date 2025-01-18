import psycopg2
import requests
from rest import get_req
from seralizer import urlgen
from validator import Metadata



class Pg_service:

    def __init__(self,dbname,user):
        self.conn = psycopg2.connect(f"dbname = {dbname} user = {user}")
        self.cursor = self.conn.cursor

    # TODO: test me
    def table_exist(self,table_name):

        self.cursor.execute("SELECT EXISTS(SELECT * \
                FROM information_schema.tables WHERE table_name = %s"), (table_name))
        return cursor.fetchone()[0]

    #Table met is a pydantic(BaseModel) class.
    def create_table(self,table_name,table_meta):
        if self.table_exist(f"{table_name}"):
            #TODO create a proper create table funtion and exists Exception
            raise Exception("Table " + table_name + " Exists")
        #maybe catch errors here if this starts causing problems.
        self.cursor.execute(f"CREATE TABLE {table_name} \
                (id serial PRIMARY KEY, num integer, data varchar);")

    def insert_to_table(self,data_list)
        self.cursor.copy_from(f,)

def main():
    conn = psycopg2.connect("dbname = testdb user=postgres")
    cur = conn.cursor

if __name__ == "__main__":
    main()



