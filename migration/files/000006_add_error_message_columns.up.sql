ALTER TABLE public.documents ADD COLUMN IF NOT EXISTS error_message text;
ALTER TABLE public.document_chunks ADD COLUMN IF NOT EXISTS error_message text;
