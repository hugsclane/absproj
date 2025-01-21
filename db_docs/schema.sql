


-- CREATE TYPE dataflowidentifier AS (
--     agencyid VARCHAR, -- is optional, reverts to all if not provided
--     dataflowid VARCHAR,
--     version int[3] -- is optional, reverts to latest version if not provided
-- )
-- CREATE TYPE datakey AS (
--     measure INT[], --list of all selected measures
--     idx INT[], --list of all selected industires
--     adjustment INT[], --list of all adjustments to dataset
--     region INT[], --list of averages?
--     frequency VARCHAR
-- )

CREATE TABLE metadata(
    id SERIAL PRIMARY KEY,
    dfid VARCHAR,
    links : VARCHAR,
    version: VARCHAR,
    agencyid: VARCHAR,
    name: VARCHAR
)

-- CREATE TABLE data_history(
--     id SERIAL PRIMARY KEY,
--     dfid dataflowidentifier,
--     dkey datakey,
--     startperiod VARCHAR(8),
--     endperiod VARCHAR(8),
--     format VARCHAR(21),
--     detail VARCHAR(14)
--     dimension VARCHAR --this takes the ID of any other dimesions to provide a cross sectional view
--     }
-- );

-- -- TODO: add get metadata history, or create a seperate metadata table.
