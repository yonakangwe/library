#!/bin/bash

# ==========================================
# PostgreSQL Restore Script
# Reads ONLY database section from config.yml
# ==========================================

CONFIG_FILE="../config.yml"
SQL_FILE="library.sql"
PG_BIN="/opt/homebrew/opt/postgresql@17/bin"
DB_HOST="127.0.0.1"

# ==============================
# Extract database section safely
# ==============================

DB_NAME=$(awk '/^database:/ {flag=1; next} /^[^ ]/ {flag=0} flag && /name:/ {print $2}' $CONFIG_FILE)
DB_USER=$(awk '/^database:/ {flag=1; next} /^[^ ]/ {flag=0} flag && /user:/ {print $2}' $CONFIG_FILE)
DB_PASS=$(awk '/^database:/ {flag=1; next} /^[^ ]/ {flag=0} flag && /password:/ {print $2}' $CONFIG_FILE)
DB_PORT=$(awk '/^database:/ {flag=1; next} /^[^ ]/ {flag=0} flag && /port:/ {print $2}' $CONFIG_FILE)

echo "Database: $DB_NAME"
echo "User: $DB_USER"
echo "Port: $DB_PORT"
echo "-----------------------------------"

# Export password
export PGPASSWORD="$DB_PASS"

# ==============================
# Check if database exists
# ==============================

DB_EXISTS=$($PG_BIN/psql -h $DB_HOST -p $DB_PORT -U $DB_USER -tAc \
"SELECT 1 FROM pg_database WHERE datname='$DB_NAME'")

if [ "$DB_EXISTS" == "1" ]; then
    echo "Database '$DB_NAME' already exists."
    read -p "Do you want to overwrite it? (yes/no): " ANSWER

    if [[ "$ANSWER" != "yes" ]]; then
        echo "Operation cancelled."
        unset PGPASSWORD
        exit 0
    fi

    echo "Terminating active connections..."
    $PG_BIN/psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -c \
    "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname='$DB_NAME';"

    echo "Dropping database..."
    $PG_BIN/dropdb -h $DB_HOST -p $DB_PORT -U $DB_USER $DB_NAME
fi

# ==============================
# Create database
# ==============================

echo "Creating database..."
$PG_BIN/createdb -h $DB_HOST -p $DB_PORT -U $DB_USER $DB_NAME

# ==============================
# Restore SQL file
# ==============================

echo "Restoring $SQL_FILE ..."
$PG_BIN/psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f $SQL_FILE

if [ $? -eq 0 ]; then
    echo " Database restored successfully!"
else
    echo " Restore failed!"
fi

unset PGPASSWORD

