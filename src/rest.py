import requests
from serializer import temp_rul_gen

def get_req(url):
    # testurl = "https://data.api.abs.gov.au/rest/data/ABS%2CCPI%2C1.1.0/1.115485.10.50.Q?startPeriod=2022&endPeriod=2024&format=jsondata&detail=full"
    # response = requests.get(testurl)
    response = requests.get(url)
    print(response.json)
    # print(testurl)
    print(temp_rul_gen())
    return response.json

def get_metadata():
    apiurl = "data.api.abs.gov.rest/dataflow/All?detail=referencestubs"
    data = get_req(url)
    data_list = data[dataflows]
    return data_list

if __name__ == "__main__":
    get_req()
