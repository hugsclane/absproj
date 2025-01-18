from psycopg import Connection
from pydantic import BaseModel

Class CRUD:

    def __init__(self, conn: Connection)
        self.conn = conn

    def _insert(self,statement: str, input: BaseModel)
