---------------- User Tables ----------------
CREATE TABLE
    IF NOT EXIST Users (
        user_id VARCHAR(64) NOT NULL,
        username VARCHAR(64) NOT NULL,
        is_active BOOL DEFAULT 0,
        PRIMARY KEY (userid)
    );

CREATE TABLE
    IF NOT EXIST UserInfo (
        user_id VARCHAR(64) NOT NULL,
        first_name VARCHAR(64) NOT NULL,
        last_name VARCHAR(64) NOT NULL,
        created_at BIGINT NOT NULL,
        PRIMARY KEY (userid)
    );