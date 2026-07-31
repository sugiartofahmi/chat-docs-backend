CREATE TABLE IF NOT EXISTS public.document_chunks (
    id          uuid        DEFAULT uuid_generate_v4() NOT NULL,
    document_id uuid                                   NOT NULL,
    chunk_index int                                    NOT NULL,
    content     text                                   NOT NULL,
    embedding   vector(1536)                           NOT NULL,
    created_at  timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at  timestamptz DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT document_chunks_pkey        PRIMARY KEY (id),
    CONSTRAINT fk_document_chunks_document FOREIGN KEY (document_id) REFERENCES public.documents (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_document_chunks_document ON public.document_chunks (document_id);
