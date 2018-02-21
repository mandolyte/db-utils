select id
from node
where id like ($1 || '%')
and id like ('%' || $2)
