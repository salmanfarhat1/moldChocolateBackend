## Running PostgreSQL Locally on macOS

To start and interact with your local PostgreSQL server, follow these steps:

### 1. Start the PostgreSQL server

Run the PostgreSQL server by specifying the data directory (adjust the path if different):

```bash
postgres -D /opt/homebrew/var/postgres

## Connect to PostgreSQL using the psql client

Open a new terminal and launch the interactive PostgreSQL shell:

```bash
psql postgres

### Useful PostgreSQL commands inside psql

```bash
psql postgres
\list show all DBs
\c DB_name list databases 
\dt show tables 

