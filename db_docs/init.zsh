
#init commands
pg_ctl init -D db -U postgres -o "-A md5 --pwfile=db_docs/pwdf.txt"



#start commands
pg_ctl -D db -l db_docs/logfile start

#login
psql postgres

##make a mistake, try --> 
#DROP SCHEMA absdata CASCADE;

#SET SEARCH_PATH TO absdata

# \i db_docs/schema.sql

#stop server
#pg_ctl -D db stop
#
