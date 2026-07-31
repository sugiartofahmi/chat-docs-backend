CREATE TABLE IF NOT EXISTS public.jobs (
    id          uuid        DEFAULT uuid_generate_v4() NOT NULL,
    job_name    varchar(255)                           NOT NULL,
    source      varchar(20)                            NOT NULL,
    payload     text                                   NOT NULL,
    status      varchar(20)                            NOT NULL DEFAULT 'pending',
    error       text,
    started_at  timestamptz,
    finished_at timestamptz,
    created_at  timestamptz DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT jobs_pkey      PRIMARY KEY (id),
    CONSTRAINT ck_jobs_source CHECK (source IN ('pool', 'scheduler')),
    CONSTRAINT ck_jobs_status CHECK (status IN ('pending', 'running', 'success', 'failed'))
);

CREATE INDEX IF NOT EXISTS idx_jobs_job_name ON public.jobs (job_name, created_at);
CREATE INDEX IF NOT EXISTS idx_jobs_pending ON public.jobs (status, created_at) WHERE status = 'pending';
