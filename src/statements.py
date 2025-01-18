__all__ = [
    "inst_metad_stmt",
    "get_metad_stmt",
    "inst_data_stmt",
    "get_data_stmt"
]

inst_metad_stmt = """
INSERT INTO metadata (
        id,
        agencyId,
        dataflowId,
        version,
        isFinal,
        name,
        names,
        description,
        descriptions,
        annotations
        ) VALUE(
        %(id)s,
        %(agencyId)s,
        %(dataflowId)s,
        %(version)s,
        %(isFinal)s,
        %(name)s,
        %(names)s,
        %(description)s,
        %(descriptions)s,
        %(annotation)s)
RETURNING id

"""
## annotations is going to break postgres
get_metad_stmt = """
SELECT * FROM metadata where id = %(id)s

"""


inst_data_stmt = """
INSERT INTO data (
        id,
        agencyId,
        dataflowId,
        version,
        index,
        adjustment,
        region,
        frequency,
        startPeriod,
        endPeriod,
        format,
        detail,
        dimensions
        ) VALUE (
        %(id)s,
        %(agencyId)s,
        %(dataflowId)s,
        %(version)s,
        %(index)s,
        %(adjustment)s,
        %(region)s,
        %(frequency)s,
        %(startPeriod)s,
        %(endPeriod)s,
        %(format)s,
        %(detail)s,
        %(dimension)s
        )
"""
get_data_stmt = """


"""
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
