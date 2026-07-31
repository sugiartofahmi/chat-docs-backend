CREATE TABLE IF NOT EXISTS public.conversations (
    id         uuid        DEFAULT uuid_generate_v4() NOT NULL,
    project_id uuid                                   NOT NULL,
    role       varchar(20)                            NOT NULL,
    content    text                                   NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT conversations_pkey       PRIMARY KEY (id),
    CONSTRAINT fk_conversations_project FOREIGN KEY (project_id) REFERENCES public.projects (id) ON DELETE CASCADE,
    CONSTRAINT ck_conversations_role    CHECK (role IN ('user', 'assistant'))
);

CREATE INDEX IF NOT EXISTS idx_conversations_project ON public.conversations (project_id, created_at);
