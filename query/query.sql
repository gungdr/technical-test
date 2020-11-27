select 
    a.ID,
    a.UserName,
    b.UserName as Parent
from `user` a
left join `user` b 
on a.Parent = b.ID
