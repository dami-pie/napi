-- +goose Up
alter table entranceLog
    modify roomNumber int not null;

alter table entranceLog
    modify entranceType int not null;

alter table entranceLog
    modify roomBuild varchar(1) not null;

alter table entranceLog
    modify userID int not null;
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
alter table entranceLog
    modify roomNumber int null;

alter table entranceLog
    modify entranceType int null;

alter table entranceLog
    modify roomBuild varchar(1) null;

alter table entranceLog
    modify userID int null;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
