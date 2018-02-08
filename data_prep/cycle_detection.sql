with dataset as
(
    select 'A' as parent, 'B' as node union all
    select 'B' as parent, 'C' as node union all
    select 'C' as parent, 'D' as node union all 
    select 'D' as parent, 'A' as node
),
hierarchy( level, parent, node, path, cyclic_flag ) as (
    select 0 as level,
    dataset.parent,
    dataset.node,
    '/' || dataset.parent || '/' || dataset.node as path,
    0 as cyclic_flag
    from dataset
    where dataset.parent = 'A'

    union all

    select
    hierarchy.level + 1 as level,
    dataset.parent,
    dataset.node,
    hierarchy.path || '/' || dataset.node as path,
    case
        when
            (length(path||dataset.node) 
            - length(replace(path|| dataset.node,dataset.node,'')))
            / length(dataset.node) = 1
        then 0
        else 1
    end as cyclic_flag
    from hierarchy
    inner join dataset
    on dataset.parent = hierarchy.node
    where 
        (length(path) 
        - length(replace(path,hierarchy.node,'')))
        / length(hierarchy.node) < 2
)
select *
from hierarchy

order by path
;