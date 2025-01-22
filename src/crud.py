from psycopg import Connection
from pydantic import BaseModel

# from src.statements import *
# from src.models import *

from statements import *
from models import *

from typing import TypeVar

T = TypeVar('T')


class CRUD:
    def __init__(self, conn: Connection):
        self.conn = conn

    def _insert(self,statement: str, input: BaseModel) -> str:
        with self.conn.cursor() as curs:
            ##TODO: this exclude annotations is a temporary measure and should be removed 
            #when validation is properly implemented
            curs.execute(statement,input.model_dump(exclude={"annotations"}))
            #checks to see if operation completed sucessfully
            return cursor.fetchone()[0]

    def _get(self,statement,model_class: type[T], **kwargs) -> T | None:
       with self.conn.cursor(row_factory=class_row(model_class)) as cursor:
           cursor.execute(statement,kwargs)
           return cursor.fetchone()

    def insert_metadata(self,input: MetadataInput) -> str:
        return self._insert(in_metad_stmt,input)

    def get_metadata(self ,id: str) -> Metadata | None:
        return self._get(get_metad_stmt,Metadata, id = id)


