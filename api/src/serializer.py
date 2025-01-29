import requests

#This serializer.py file is used to parse from the ABS api db into a rest URL
#TODO:
#consider renaming since serializer might be better used for the local db.

#data flow identifer in the form
#{agencyId},{dataflowId},{version}

def dataflowserializer(agencyId ,dataflowId ,version) -> str:
    fmt = '{fmt:}'
    data = agencyId + fmt + dataflowId + fmt + version
    return data.format(fmt = ",")

def datakeyserializer(measure,index,adjustment,region,frequency) -> str:
    fmt = '{fmt:}'
    data = measure + fmt + index + fmt + adjustment + fmt + region + fmt + frequency
    return data.format(fmt = ".")

def queryserializer(startPeriod="",endPeriod="",format="",detail="",dimension=""):
    fmt = '{fmt:}'
    if startPeriod != "":
        startPeriod = "startPeriod="+startPeriod
    if endPeriod != "":
        endPeriod = fmt + "endPeriod=" + endPeriod
    if format != "":
        format = fmt + "format=" + format
    if detail != "":
        detail = fmt + "detail=" + detail
    if dimension != "":
        dimension= fmt + "dimension=" + dimension
    data = "?" +startPeriod + endPeriod + format + detail + dimension
    return data.format(fmt = "&")

def urlGen(api_url,dataflow,datakey,query):
    # print(dataflow)
    dataflow = dataflow.replace(",","%2C")
    fmt = '{fmt:}'
    url = api_url + fmt + dataflow + fmt + datakey + query
    return url.format(fmt = "/")

def temp_rul_gen():
    #dataflowserializer
    api_url = "https://data.api.abs.gov.au/rest/data"
    agencyId = "ABS"
    dataflowId = "CPI"
    version = "1.1.0"
    dataflow = dataflowserializer(agencyId,dataflowId,version)

    #datakeyserializer
    measure = "1"
    index = "115485"
    adjustment_type = "10"
    region = "50"
    frequency = "Q"
    datakey = datakeyserializer(measure,index,adjustment_type,region,frequency)

    #queryserializer
    startPeriod = "2015"
    endPeriod = "2024"
    format = "jsondata"
    detail = "full"
    query = queryserializer(startPeriod,endPeriod,format,detail)
    url = urlGen(api_url,dataflow,datakey,query)
    # print(url)
    return url

if __name__ == "__main__":

    #tests

    #dataflowserializer
    api_url = "data.api.abs.gov.au/rest/data"
    agencyId = "ABS"
    dataflowId = "CPI"
    version = "1.1.0"
    dataflow = dataflowserializer(agencyId,dataflowId,version)

    #datakeyserializer
    measure = "1"
    index = "11548"
    adjustment_type = "10"
    region = "50"
    frequency = "Q"
    datakey = datakeyserializer(measure,index,adjustment_type,region,frequency)

    #queryserializer
    startPeriod = "2015"
    endPeriod = "2024"
    format = "jsondata"
    detail = "full"
    query = queryserializer(startPeriod,endPeriod,format,detail)
    url = urlGen(api_url,dataflow,datakey,query)
    print(url)
