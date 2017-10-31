## Agenda 命令设计

>agenda help [option]

列出命令说明(可选是否列出具体功能的说明)。

>agenda register -uUserName --password pass -email=a@xxx.com -phone=phoneNum

用户注册，如果用户名已被使用，返回错误信息；如果注册成功，返回成功信息。

>agenda login -uUserName --password pass

用户登录，登录失败返回失败原因;登录成功，返回成功信息，并列出可选操作。

>agenda logout

用户退出登录，返回成功信息并列出可选操作。

>agenda -lUser

已登录用户查询已注册用户信息（用户名、邮箱、电话）

>agenda delete

已登录用户注销帐号，操作成功返回成功信息；否则，返回失败信息。若成功，删除一切与该用户的信息。

>agenda mkmeeting --title Name --participator user1 [user2 ....] --stime start --etime end

成功，则返回成功信息及注册信息；失败，则返回失败原因。

>agenda meetingadd --participator user1 [user2 ...]

成功，则返回新增参与者信息；失败，返回失败原因。

> agenda meetingdel --participator user1 [user2 ...]

成功，则返回成功信息（如果会议因为删除参与者而删除，返回额外信息）；失败，返回失败原因。

>agenda querymeeting -stime start -etime end

已登录用户查询自己的会议议程。

>agenda meetingdel --title Name

已登录用户删除会议。

>agenda meetingout --ttile Name

已登录用户退出自己参与的某一会议，若因此造成会议参与者为0,则会议被删除。

>agenda meetingclear

已登录用户清空自己发起的所有会议。
