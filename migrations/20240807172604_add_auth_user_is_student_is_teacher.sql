-- +goose Up
ALTER TABLE public.auth_user
    ADD COLUMN is_student BOOLEAN NOT NULL DEFAULT true,
    ADD COLUMN is_teacher BOOLEAN NOT NULL DEFAULT false;
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE public.auth_user
    DROP COLUMN is_student,
    DROP COLUMN is_teacher;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
