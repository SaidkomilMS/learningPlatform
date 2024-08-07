-- +goose Up
ALTER TABLE public.auth_user ADD COLUMN created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now();
ALTER TABLE public.auth_user ADD COLUMN updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now();
ALTER TABLE public.auth_user ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE;
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE public.auth_user DROP COLUMN created_at;
ALTER TABLE public.auth_user DROP COLUMN updated_at;
ALTER TABLE public.auth_user DROP COLUMN deleted_at;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
