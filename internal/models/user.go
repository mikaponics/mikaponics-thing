package models  // github.com/mikaponics/mikapod-iam/internal/models


import (
    "database/sql"
    "fmt"

    "github.com/mikaponics/mikaponics-thing/internal/utils"
)


type User struct {
    TenantId          int64          `db:"tenant_id"`
    TenantSchema      sql.NullString `db:"tenant_schema"`
    Id                int64          `db:"id"`
    FirstName         sql.NullString `db:"first_name"`
    LastName          sql.NullString `db:"last_name"`
    PasswordHash      sql.NullString `db:"password_hash"`
    Email             sql.NullString `db:"email"`
    RoleId            uint8          `db:"role_id"`
}


// Special Thanks:
// * https://jmoiron.github.io/sqlx/
// * http://wysocki.in/golang-sqlx/

/**
 *  Function will create the `users` table in the database.
 */
func (dal *DataAccessLayer) CreateUserTable(dropExistingTable bool) {
    if dropExistingTable {
        drop_stmt := "DROP TABLE users;"
        results, err := dal.db.Exec(drop_stmt)
        if err != nil {
            fmt.Println("User Model:", results, err)
        }
    }

    // Special thanks:
    // * http://www.postgresqltutorial.com/postgresql-create-table/
    // * https://www.postgresql.org/docs/9.5/datatype.html

    stmt := `CREATE TABLE users (
        tenant_id bigint NOT NULL,
        tenant_schema VARCHAR (127) NOT NULL,
        id bigserial PRIMARY KEY,
        first_name VARCHAR (50) NOT NULL,
        last_name VARCHAR (50) NOT NULL,
        password_hash VARCHAR (511) NOT NULL,
        email VARCHAR (255) UNIQUE NOT NULL,
        role_id INT NOT NULL
    );`
    results, err := dal.db.Exec(stmt)
    if err != nil {
        fmt.Println("User Model", results, err)
    }
    return
}


/**
 *  Function will return the `user` struct if it exists in the database or
 *  return an error.
 */
func (dal *DataAccessLayer) FindUserByEmail(email string) (*User, error) {
    user := User{} // The struct which will be populated from the database.

    // DEVELOPERS NOTE:
    // (1) Lookup the user based on the email.
    // (2) PostgreSQL uses an enumerated $1, $2, etc bindvar syntax
    err := dal.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)

    // Handling non existing item
    if err == sql.ErrNoRows {
        return nil, nil
    } else if err != nil {
        return nil, err
    }

    return &user, nil
}


/**
 *  Function will create a user, if validation passess, and reutrns the `user`
 *  struct else returns the error.
 */
func (dal *DataAccessLayer) CreateUser(email string, firstName string, lastName string, password string, tenantId int64, tenantSchema string, roleId int64) (*User, error) {
    // Step 1: Hash our user's password for added security or error on any condition.
    passwordHash, err := utils.HashPassword(password)
    if err != nil {
        return nil, err
    }

    // Step 2: Generate SQL statement for creating a new `user` in `postgres`.
    statement := `INSERT INTO users (email, first_name, last_name, password_hash, tenant_id, tenant_schema, role_id) VALUES ($1, $2, $3, $4, $5, $6, $7)`

    // Step 3: Execute our SQL statement and either return our new user or
    //         our error.
    _, err = dal.db.Exec(statement, email, firstName, lastName, passwordHash, tenantId, tenantSchema, roleId)
    if err != nil {
        return nil, err
    }
    return dal.FindUserByEmail(email)
}
