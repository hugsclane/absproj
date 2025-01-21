
#init commands
pg_ctl init -D db -U postgres -o "-A md5 --pwfile=db_docs/pwdf.txt"



#start commands
pg_ctl -D db -l db_docs/logfile start

#login
psql postgres
