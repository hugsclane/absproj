import json
import os

with open('metadata/reference_stubs.json','r') as file:
    dfile = json.load(file)
# print(os.getcwd())
kdict = []

for k ,v in dfile.items():
    if k == "data":
        for key, value in v.items():

            for i,j in enumerate(value):
                if i == 1096:
                    print(str(j) + "\n")
                if i == 1097:
                    print(j)
                    break

        #         kdict.append(i['name'])
            # kdict[key] = value
            # break
for i in kdict:
    print(i + "\n")

