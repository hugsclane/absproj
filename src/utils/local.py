# from rest import get_json
import json

# ensure that these two functions are called from the root directory of this project for consistency

__all__ = [
    "pwfromfile",
    "of_json"
        ]
def pwfromfile(dir):
    # "db_docs/pwdf.txt"
    with open(dir,'r', encoding = "UTF-8") as file:
        val = file.read()
        file.close()
        return val.strip("\n")

def of_json(dir):
    #"metadata/reference_stubs.json"
    with open(dir, 'r') as file:
            data = json.loads(file.read())
            file.close()
            return data


