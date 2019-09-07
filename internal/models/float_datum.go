package models  // github.com/mikaponics/mikapod-iam/internal/models


import (
    "database/sql"
    "fmt"

    // "github.com/mikaponics/mikaponics-thing/internal/utils"
)


type TimeSeriesFloatDatum struct {
    Id                int64          `db:"id"`
    TenantId          int64          `db:"tenant_id"`
    SensorId          int64          `db:"sensor_id"`
    value             float64        `db:"value"`
    timestamp         int64          `db:"timestamp"`
}

/*
 * typeId
 * (1) Temperature
 */

// Special Thanks:
// * https://jmoiron.github.io/sqlx/
// * http://wysocki.in/golang-sqlx/

/**
 *  Function will create the `time series float data` table in the database.
 */
func (dal *DataAccessLayer) CreateTimeSeriesFloatDatumTable(dropExistingTable bool) {
    if dropExistingTable {
        drop_stmt := "DROP TABLE data;"
        results, err := dal.db.Exec(drop_stmt)
        if err != nil {
            fmt.Println("TimeSeriesFloatDatum Model:", results, err)
        }
    }

    // Special thanks:
    // * http://www.postgresqltutorial.com/postgresql-create-table/
    // * https://www.postgresql.org/docs/9.5/datatype.html

    stmt := `CREATE TABLE data (
        id BIGSERIAL PRIMARY KEY,
        tenant_id BIGINT NOT NULL,
        sensor_id BIGINT NOT NULL,
        value DECIMAL NULL,
        timestamp BIGINT NOT NULL
    );`
    results, err := dal.db.Exec(stmt)
    if err != nil {
        fmt.Println("TimeSeriesFloatDatum table dropped with error", results, err)
    } else {
        fmt.Println("TimeSeriesFloatDatum table dropped and re-created")
    }
    return
}


/**
 *  Function will return the `sensor` struct if it exists in the database or
 *  return an error.
 */
func (dal *DataAccessLayer) GetTimeSeriesFloatDatumByTenantIdAndTimestamp(tenantId int64, timestamp int64) (*TimeSeriesFloatDatum, error) {
    datum := TimeSeriesFloatDatum{} // The struct which will be populated from the database.

    // DEVELOPERS NOTE:
    // (1) Lookup the datum based on the email.
    // (2) PostgreSQL uses an enumerated $1, $2, etc bindvar syntax
    err := dal.db.Get(&datum, "SELECT * FROM data WHERE tenant_id = $1 AND timestamp = $2 LIMIT 1", tenantId, timestamp)

    // Handling non existing item
    if err == sql.ErrNoRows {
        return nil, nil
    } else if err != nil {
        return nil, err
    }

    return &datum, nil
}


/**
 *  Function will create a datum, if validation passess, and reutrns the `datum`
 *  struct else returns the error.
 */
func (dal *DataAccessLayer) CreateTimeSeriesFloatDatum(tenantId int64, sensorId int64, value float64, timestamp int64) (*TimeSeriesFloatDatum, error) {
    // Step 1: Generate SQL statement for creating a new `user` in `postgres`.
    statement := `INSERT INTO data (tenant_id, sensor_id, value, timestamp) VALUES ($1, $2, $3, $4)`

    // Step 2: Execute our SQL statement and either return our new user or
    //         our error.
    _, err := dal.db.Exec(statement, tenantId, sensorId, value, timestamp)
    if err != nil {
        return nil, err
    }

    // Step 3:
    return dal.GetTimeSeriesFloatDatumByTenantIdAndTimestamp(tenantId, timestamp)
}
