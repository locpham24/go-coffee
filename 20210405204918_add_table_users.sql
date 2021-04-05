-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	id SERIAL NOT NULL PRIMARY KEY,
	phone_number varchar(50) NOT NULL,
	password TEXT,
	created_by int4 NOT NULL,
	created_at timestamptz NOT NULL DEFAULT now(),
	updated_by int4 NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
