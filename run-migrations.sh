source ./migrations.env && \
migrate \
  -database ${DB_URL} \
  -path ./migrations \
  up 