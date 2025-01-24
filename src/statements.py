__all__ = [
    "in_metad_stmt",
    "get_metad_stmt",
    "in_data_stmt",
    "get_data_stmt"
]

in_metad_stmt = """
INSERT INTO absdata.metadata (
        id,
        agencyId,
        dataflowId,
        version,
        isFinal,
        name,
        description
        ) VALUES (
        %(id)s,
        %(agencyId)s,
        %(dataflowId)s,
        %(version)s,
        %(isFinal)s,
        %(name)s,
        %(description)s
        )
RETURNING id

"""
## annotations is going to break postgres
get_metad_stmt = """
SELECT * FROM absdata.metadata where id = %(id)s
"""


in_data_stmt = """
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
