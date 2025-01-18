import json
import os

with open('metadata/abs_dataflow.json','r') as file:
    dfile = json.load(file)
# print(os.getcwd())
kdict = []

for k ,v in dfile.items():
    if k == "data":
        for key, value in v.items():
            for i in value:
                kdict.append(i['name'])  
            # kdict[key] = value
            # break
for i in kdict:
    print(i + "\n")
        
