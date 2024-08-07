-- +goose Up
ALTER TABLE public.auth_user
    ALTER COLUMN is_superuser SET DEFAULT false,
    ALTER COLUMN is_staff SET DEFAULT false,
    ALTER COLUMN is_active SET DEFAULT true,
    ALTER COLUMN is_test SET DEFAULT false,
    ALTER COLUMN date_joined SET DEFAULT now();
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE public.auth_user
    ALTER COLUMN is_superuser DROP DEFAULT,
    ALTER COLUMN is_staff DROP DEFAULT,
    ALTER COLUMN is_active DROP DEFAULT,
    ALTER COLUMN is_test DROP DEFAULT,
    ALTER COLUMN date_joined DROP DEFAULT;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
