


-- CREATE TYPE dataflowidentifier AS (
--     agencyid TEXT, -- is optional, reverts to all if not provided
--     dataflowid TEXT,
--     version int[3] -- is optional, reverts to latest version if not provided
-- )
-- CREATE TYPE datakey AS (
--     measure INT[], --list of all selected measures
--     idx INT[], --list of all selected industires
--     adjustment INT[], --list of all adjustments to dataset
--     region INT[], --list of averages?
--     frequency TEXT
-- )
-- TODO format
CREATE SCHEMA absdata;

-- CREATE TYPE texttextpair AS (k TEXT, V TEXT);

CREATE TABLE absdata.metadata(
    id CHAR(32) PRIMARY KEY,
    dataflowid TEXT,
    version TEXT,
    isfinal BOOLEAN,
    agencyid TEXT,
    description TEXT,
    name TEXT,
    UNIQUE (dataflowid)
);

CREATE TABLE absdata.datablob(
  id CHAR(32) PRIMARY KEY,
  dataflowid TEXT,
  version TEXT,
  name TEXT,
  blob,
  UNIQUE(dataflowid)
)

--Excluding annotations

-- CREATE TABLE data_history(

--     id SERIAL PRIMARY KEY,
--     dfid dataflowidentifier,
--     dkey datakey,
--     startperiod TEXT(8),
--     endperiod TEXT(8),
--     format TEXT(21),
--     detail TEXT(14)
--     dimension TEXT --this takes the ID of any other dimesions to provide a cross sectional view
--     }
-- );

-- -- TODO add get metadata history, or create a seperate metadata table.
