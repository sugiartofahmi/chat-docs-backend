CREATE TABLE IF NOT EXISTS public.projects (
    id         uuid        DEFAULT uuid_generate_v4() NOT NULL,
    name       varchar(255)                           NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamptz,
    CONSTRAINT projects_pkey PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_projects_name ON public.projects (name) WHERE deleted_at IS NULL;
