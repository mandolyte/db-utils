CREATE TABLE if not exists  public.edge (
    id text primary key,
    from_id text not null,
    to_id text not null
)
