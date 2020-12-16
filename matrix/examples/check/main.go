package main

import (
	log "github.com/sirupsen/logrus"
	matrix "github.com/skeptycal/matrix"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.Info("Go MySQL connection check")

	matrix.Check()

}

// Notes:
/* Matrix - an organizational map based on the Eisenhower Matrix.
mysql - https://www.tutorialspoint.com/mysql/mysql-create-tables.htm

// -----------------------------------------------------------> todo - list of steps to complete

// ➜  install and test mysql 8.x
     brew install mysql
     brew services start mysql
// We've installed your MySQL database without a root password. To secure it run:
     mysql_secure_installation
     brew services restart mysql
     zbrew services list
// MySQL is configured to only allow connections from localhost by default
// To connect run:
     mysql -uroot
*/

// // -----------------------------------------------------------> mysqlsh
/*
The following modules and objects are ready for use when the shell starts:

 - dba     Used for InnoDB cluster administration.
 - mysql   Support for connecting to MySQL servers using the classic MySQL
           protocol.
 - mysqlx  Used to work with X Protocol sessions using the MySQL X DevAPI.
 - plugins Plugin to manage MySQL Shell plugins
 - session Represents the currently open MySQL session.
 - shell   Gives access to general purpose functions and properties.
 - util    Global object that groups miscellaneous tools like upgrade checker
           and JSON import.

// using mysqlsh:

\status
    MySQL Shell version 8.0.22

    Not Connected.

// switch to python mode (personal preference)
\py

// // -----------------------------------------------------------> cli connection
\c root@localhost
// or \connect root@localhost:33060
// (save password if desired)

\status
// server response - ymmv
    MySQL Shell version 8.0.22

    Connection Id:                19
    Default schema:
    Current schema:
    Current user:                 root@localhost
    SSL:                          Cipher in use: TLS_AES_256_GCM_SHA384 TLSv1.3
    Using delimiter:              ;
    Server version:               8.0.22 MySQL Community Server - GPL
    Protocol version:             X protocol
    Client library:               8.0.22
    Connection:                   localhost via TCP/IP
    TCP port:                     33060
    Server characterset:          utf8mb4
    Schema characterset:          utf8mb4
    Client characterset:          utf8mb4
    Conn. characterset:           utf8mb4
    Result characterset:          utf8mb4
    Compression:                  Enabled (DEFLATE_STREAM)
    Uptime:                       1 day 5 hours 42 min 58.0000 sec

// if using \connect root@localhost:3306 (prefer 33060)
    MySQL Shell version 8.0.22

    Connection Id:                21
    Current schema:
    Current user:                 root@localhost
    SSL:                          Cipher in use: TLS_AES_256_GCM_SHA384 TLSv1.3
    Using delimiter:              ;
    Server version:               8.0.22 MySQL Community Server - GPL
    Protocol version:             Classic 10
    Client library:               8.0.22
    Connection:                   localhost via TCP/IP
    TCP port:                     3306
    Server characterset:          utf8mb4
    Schema characterset:          utf8mb4
    Client characterset:          utf8mb4
    Conn. characterset:           utf8mb4
    Result characterset:          utf8mb4
    Compression:                  Disabled
    Uptime:                       1 day 5 hours 49 min 10.0000 sec


// execute code from a script (CAREFUL!!)
\sql

\source /tmp/mydata.sql

SHOW DATABASES;
SHOW SCHEMAS;
SHOW DATABASES LIKE 'open%';
// The percent sign (%) means zero, one, or multiple characters.
USE sys;
SHOW TABLES;
    +-----------------------------------------------+
    | Tables_in_sys                                 |
    +-----------------------------------------------+
    | host_summary                                  |
    | host_summary_by_file_io
    ...

SHOW FULL TABLES;
SHOW TABLES FROM sys;

SELECT – selects data from the database.
DELETE – removes data from the database.
UPDATE – overwrites data in the database.
CREATE DATABASE – creates a new database.
INSERT INTO – uploads new data into the database.
CREATE TABLE – creates a new table.
ALTER DATABASE – changes the attributes, files or filegroups of the database.
CREATE INDEX – creates an index, or a search key.
DROP INDEX – deletes an index.
ALTER TABLE – changes the attributes or entries of a table.


\js
session
    <Session:root@localhost:33060>
\status



*/

//// Data:
/*
- structure to store actions
- structure to store calendar days
- structure to store resources
- structure to store ideas
- structure to store relationship graphs

// Major Parts:
- mysql database
- go link to mysql database
- frontend templates

// CRUD database
Create -

Read -

Update -

Delete -

// API funcionality

// API Documentation


*/
