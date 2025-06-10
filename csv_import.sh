#!/bin/bash

until pg_isready -h db -p 5432; do echo 'Waiting for Postgres...'; sleep 2; done;
ROWS=$(psql -h db -U ismael -d ilumeo -t -c "SELECT count(*) FROM historico;");
if [ "$ROWS" -eq "0" ]; then
  echo 'Importing CSV...';
  psql -h db -U ismael -d ilumeo -c "\COPY historico(id, origin, response_status_id) FROM '/csv/populate_historico.csv' DELIMITER ',' CSV;";
else
  echo 'Skipping import, historico table already has data.';
fi
