### Intent
This project is a work in progress to improve my skills in a variaty of technical areas, and in project management.

For now this is a solo project, so my methodology for working through the tech stack is going to be an iterative one. I will make changes in an environment until I'm confident that I understand tehcnology that I am working with, and I will leave sub-repo's incomplete unitl I have a better grasp on the ultimate structure of the project.

The API folder contains inital exploration into pydantic and pyscopg, I intended to use flaks initally but have since decided to use Go to build the webserver.

Astro-Svelte is a framework I'm familiar with and I don't want to focus on frontend dev too much. If my priorties chagne, I'll likely invest time into learning html

### Strucutre (Subject to change)

 - Python for the modeling application used to express data from the ABS db
 - Astro-Svelte website (might make a HTMX website in parrallel)
 - Go webserver



### Requirements (incomplete)
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

Install GNUmake

Run make postgres
Run make redis

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

chrisjpalmer for guidance on golang and webserver setup
https://github.com/chrisjpalmer
