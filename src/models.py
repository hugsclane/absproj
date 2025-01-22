from pydantic import BaseModel, computed_field
from typing import Optional, Dict
import psycopg

#TODO: fix this up ref schema.sql.
class Search(BaseModel):
    id: int
    # dataflowId
    agencyId: Optional[str] = None #Default All
    dataflowId: str
    version: Optional[str] = None #Default Latest
    # datakey
    index: int
    adjustment: str
    region: int
    #optional data
    frequency: str
    startPeriod: str
    endPeriod: str
    format: str
    detail: str
    dimension: str

    # @validator("agencyId")
    # def validate_agencyId(cls,value):
    #     agency_list = [
    #             "SDMX",
    #             "ABS",
    #             "All"
    #             ]
    #     if value not in agency_list:
    #         raise ValueError(f"invalid agencyID :{value}")
    #     return value
    # @validator("dataflowId")
    # def dataflowId(cls,value):

class Metadata(BaseModel):
    id: str
    # dataflowId
    agencyid: Optional[str] = None #default all
    dataflowid: str
    version: Optional[str] = None #default latest
    links: Dict[str,str]
    isFinal : bool
    name : str
    names : Dict[str,str]
    # annotations: list

class MetadataInput(BaseModel):

    agencyid: Optional[str] = None #default all
    dataflowid: str
    links: Dict[str,str]
    version: Optional[str] = None #default latest
    isFinal : bool
    name : str
    names: Dict[str,str]
    # annotations: list

    @computed_field
    def id(self) -> str:
        return uuid.hex

if __name__ == "__main__":
    pass
