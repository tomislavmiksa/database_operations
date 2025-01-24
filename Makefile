# postgres database variables
container_image = postgres
container_name = test-postgres
container_port = 5432
db_user = postgres
db_password = postgres
db_name = simple_bank

# start a postgress database as a container in podman
# will as well create the simple_bank database
# creates the basic schema we need
start:
	podman run --name $(container_name) -p $(container_port):$(container_port) -e POSTGRES_PASSWORD=$(db_password) -d $(container_image)
	sleep 5
	podman exec -it $(container_name) createdb -U $(db_user) --owner=$(db_user) $(db_name)
	podman cp ./sql_schema/postgre_schema_sample.sql $(container_name):/tmp/postgre_schema_sample.sql
	podman exec -u postgres $(container_name) psql postgres postgres -f /tmp/postgre_schema_sample.sql

# stop a postgress database as a container in podman
status:
	podman ps -a

# delete a postgress database as a container in podman
clean:
	podman exec -it $(container_name) dropdb -U $(db_user) $(db_name)
	podman stop $(container_name)
	sleep 5
	podman rm $(container_name)

# connect to the database
# useful command list for the 1st hand
# - list databases: \l
# - connect to database: \c <db_name>
# - list tables: \dt
# - list table content: select * from <db_name>.<table_name> LIMIT 10;
connect:
	podman exec -it $(container_name) psql -U $(db_user)
