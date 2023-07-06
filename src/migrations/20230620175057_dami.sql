-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE userGroup(
    userGroupId int PRIMARY KEY,
    userGroupDesc varchar(50) NOT NULL
);

CREATE TABLE userdata(
    userID int auto_increment PRIMARY KEY,
    userEmail varchar(150) NOT NULL unique,
    userGroupId int,
    FOREIGN KEY (userGroupId) REFERENCES userGroup(userGroupId)
);

CREATE TABLE zones(
    zoneId int PRIMARY KEY,
    zoneDesc varchar(50) NOT NULL
);

CREATE TABLE accessPermit(
    accessPermitDesc varchar(50) NOT NULL,
    userGroupId int,
    zoneId int,
    FOREIGN KEY (userGroupId) REFERENCES userGroup(userGroupId),
    FOREIGN KEY (zoneId) REFERENCES zones(zoneId)
);

CREATE TABLE room(
    roomNumber int,
    zoneId int,
    roomBuild varchar(1),
    CONSTRAINT roomKey PRIMARY KEY (roomNumber, roomBuild, zoneId),
    FOREIGN KEY (zoneId) REFERENCES zones(zoneId)
);

CREATE TABLE userCard(
    cardHash varchar(50) PRIMARY KEY,
    cardEnable int NOT NULL,
    userID int,
    FOREIGN KEY (userID) REFERENCES userdata(userID)
);

CREATE TABLE entranceLog(
    roomNumber int,
    entranceType int,
    roomBuild varchar(1),
    userID int,
    dateHour DATE NOT NULL,
    cardHash varchar(50),
    FOREIGN KEY (userID) REFERENCES userdata(userID),
    CONSTRAINT roomKey FOREIGN KEY (roomNumber, roomBuild) REFERENCES room (roomNumber, roomBuild),
    FOREIGN KEY (cardHash) REFERENCES userCard(cardHash)
);

ALTER TABLE userCard ALTER cardEnable SET DEFAULT 0;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

--down
DROP TABLE entranceLog;
DROP TABLE userCard;
DROP TABLE room;
DROP TABLE accessPermit;
DROP TABLE zones;
DROP TABLE userdata;
DROP TABLE userGroup;