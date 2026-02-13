#!/bin/bash

# ==========================================
# PostgreSQL Restore Script (Auto-Detect v17)
# Reads ONLY database section from config.yml
# ==========================================

CONFIG_FILE="../config.yml"
SQL_FILE="library.sql"
DB_HOST="127.0.0.1"

echo "-----------------------------------"
echo "PostgreSQL Restore Script"
echo "-----------------------------------"

# ==============================
# Check required files
# ==============================

if [ ! -f "$CONFIG_FILE" ]; then
    echo "Error: config.yml not found!"
    exit 1
fi

if [ ! -f "$SQL_FILE" ]; then
    echo "Error: $SQL_FILE not found!"
    exit 1
fi

# ==============================
# Auto-detect PostgreSQL 17
# ==============================

if brew list postgresql@17 &>/dev/null; then
    PG_BIN="$(brew --prefix postgresql@17)/bin"
elif command -v psql &>/dev/null; then
    PG_BIN="$(dirname $(command -v psql))"
else
    echo "Error: PostgreSQL client not found."
    exit 1
fi

PG_VERSION=$($PG_BIN/psql --version | awk '{print $3}' | cut -d. -f1)

if [ "$PG_VERSION" != "17" ]; then
    echo "Error: PostgreSQL 17 required. Found version $PG_VERSION"
    exit 1
fi

echo "Using PostgreSQL binaries from: $PG_BIN"
echo "Detected version: $PG_VERSION"
echo "-----------------------------------"

# ==============================
# Extract database section safely
# ==============================

DB_NAME=$(awk '/^database:/ {flag=1; next} /^[^ ]/ {flag=0} flag && /name:/ {print $2}' $CONFIG_FILE)
DB_USER=$(awk '/^database:/ {flag=1; next} /^[^ ]/ {flag=0} flag && /user:/ {print $2}' $CONFIG_FILE)
DB_PASS=$(awk '/^database:/ {flag=1; next} /^[^ ]/ {flag=0} flag && /password:/ {print $2}' $CONFIG_FILE)
DB_PORT=$(awk '/^database:/ {flag=1; next} /^[^ ]/ {flag=0} flag && /port:/ {print $2}' $CONFIG_FILE)

DB_NAME="library"
DB_USER="postgres"
DB_PORT="5432"

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

if [ $? -ne 0 ]; then
    echo "Failed to create database!"
    unset PGPASSWORD
    exit 1
fi

# ==============================
# Restore SQL file
# ==============================

echo "Restoring $SQL_FILE ..."
$PG_BIN/psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f $SQL_FILE

if [ $? -eq 0 ]; then
    echo "Database restored successfully!"
else
    echo "Restore failed!"
fi

unset PGPASSWORD
echo "-----------------------------------"

