from pydantic import BaseModel
from typing import Optional
import psycopg


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
    isFinal : bool
    name : str
    names : dict
    description: str
    descriptions: dict
    annotations: list

class MetadataInput(BaseModel):

    agencyid: Optional[str] = None #default all
    dataflowid: str
    version: Optional[str] = None #default latest
    isFinal : bool
    name : str
    description: str
    annotations: list

if __name__ == "__main__":
    pass
