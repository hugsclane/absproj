curl 'https://data.api.abs.gov.au/rest/data/ABS,RES_DWELL/1.1+1GSYD+1RNSW.Q?detail=Full&startPeriod=2020-Q1&endPeriod=2020-Q2&format=jsondata' \
| jq > ./scripts/out.json
