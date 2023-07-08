-- +goose Up
alter table entranceLog
    drop foreign key entranceLog_ibfk_1;

alter table entranceLog
    drop foreign key entranceLog_ibfk_2;
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
alter table entranceLog
    add foreign key (userID) references userdata(userID);

alter table entranceLog
    add foreign key (cardHash) references userCard(cardHash);

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
