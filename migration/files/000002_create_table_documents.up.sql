CREATE TABLE IF NOT EXISTS public.documents (
    id         uuid        DEFAULT uuid_generate_v4() NOT NULL,
    project_id uuid                                   NOT NULL,
    filename   varchar(255)                           NOT NULL,
    status     varchar(50) DEFAULT 'processing'        NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamptz,
    CONSTRAINT documents_pkey       PRIMARY KEY (id),
    CONSTRAINT fk_documents_project FOREIGN KEY (project_id) REFERENCES public.projects (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_documents_project ON public.documents (project_id);
