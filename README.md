### Intent
This project is a work in progress to improve my skills in a variaty of technical areas, and in project management. 

### Requirements
GNUmake
Dagger
Docker


# Project Description
A website providing access to ABS data and some modeling tools with that data.

## Technologies
- python
    - pydantic
- postgres
- poetry
- Astro

#### Webservice install

Install a version of go higher than go 1.21.

Install docker

Install dagger

##### Makefile

- `dagger call redis up`
- `dagger call postgres up`
- `dagger call rebrow up`

- `make run`: runs the project locally
- `make docker-build`: builds the project in docker
- `make docker-run`: runs the project in docker


### References

##### The api reference used to as a reference for structure.
https://api.gov.au/assets/APIs/abs/DataAPI.openapi.html


##### Thanks
dakivara pydantic_with_psycopg tutorial
https://github.com/dakivara/pydantic_with_psycopg

chrisjpalmer for guidance on golang
https://github.com/chrisjpalmer

