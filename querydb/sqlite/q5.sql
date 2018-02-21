select id
from node
where id like (? || '%')
and id like ('%' || ?)