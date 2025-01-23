from uuid import uuid4
from pydantic import BaseModel, computed_field
from typing import Optional, Dict
import psycopg


__all__ = [
        "Search",
        "Metadata",
        "MetadataInput"
        ]


#TODO: fix this up ref schema.sql.
class Search(BaseModel):
    id: int
    # dataflowId
    agencyID: str #Default All
    dataflowId: str
    version: str #Default Latest
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

    # @validator("agencyID")
    # def validate_agencyID(cls,value):
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
##### old method for the datafile "abs_dataflow.json" for simplicity swapped to reference_stubs.sjon
# class Metadata(BaseModel):
#     id: str
#     # dataflowId
#     agencyID: str #default all
#     dataflowId: str
#     version: str #default latest
#     links: Dict[str,str]
#     isFinal : bool
#     name : str
#     names : Dict[str,str]
#     # annotations: list

# class MetadataInput(BaseModel):

#     agencyID: str #default all
#     dataflowId: str
#     links: Dict[str,str]
#     version: str #default latest
#     isFinal : bool
#     name : str
#     names: Dict[str,str]
#     # annotations: list

class Metadata(BaseModel):
    id: str

    agencyId: str #default all
    dataflowId: str
    version: str #default latest
    isFinal : bool
    name : str
    description: Optional[str] = None
    # annotations: list

class MetadataInput(BaseModel):

    agencyId: str #default all
    dataflowId: str
    version: str #default latest
    isFinal : bool
    name : str
    description: Optional[str] = None
    # annotations: list


    @computed_field
    def id(self) -> str:
        return uuid4().hex

if __name__ == "__main__":
    pass
