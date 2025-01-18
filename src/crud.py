from psycopg import Connection
from pydantic import BaseModel

Class CRUD:

    def __init__(self, conn: Connection)
        self.conn = conn

    def _insert(self,statement: str, input: BaseModel) -> str:
        with self.conn.cursor() as curs:
            curs.execute(statement,input.model_dump())
            #checks to see if operation completed sucessfully
            return cursor.fetchone()[0]

   def _get(self,statement,model_class: type[T], **kwargs):
       with self.conn.cursor(row_factory=class_row(model_class)) as cursor:
           cursor.execute(statement,kwargs)
           return cursor.fetchone()

    def insert_metadata(self,input: MetadataInput) -> str:
        return self._insert()
