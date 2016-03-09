package main

import "services/user"

func CreateUsers() (err error) {
	//项目管理员
	err = createUser("唐霞", "chinaruntx@chinarun.com", "123456", 2, []user.RoleType{user.ProjectManager})
	err = createUser("侯佳", "houjia@chinarun.com", "123456", 1, []user.RoleType{user.ProjectManager})

	//业务人员
	//SH
	//err = createUser("陈雷", "chen@chinarun.com", "123456", 1, []user.RoleType{user.BussinessMen})
	err = createUser("刘明忠", "liumingzhong@chinarun.com", "123456", 1, []user.RoleType{user.BussinessMen})
	err = createUser("李斌", "libin@chinarun.com", "123456", 1, []user.RoleType{user.BussinessMen})
	err = createUser("段司政", "duansizheng@chinarun.com", "123456", 1, []user.RoleType{user.BussinessMen})
	err = createUser("高琼如", "gaoqiongru@chinarun.com", "123456", 1, []user.RoleType{user.BussinessMen})
	err = createUser("马欢", "mahuan@chinarun.com", "123456", 1, []user.RoleType{user.BussinessMen})

	//YZ
	err = createUser("罗婷", "runffting@runff.com", "123456", 2, []user.RoleType{user.BussinessMen})
	err = createUser("朱茜", "runffzx@runff.com", "123456", 2, []user.RoleType{user.BussinessMen})
	err = createUser("吴云", "runffwy@runff.com", "123456", 2, []user.RoleType{user.BussinessMen})

	//制作人员--技术
	//SH
	err = createUser("钱高明", "qiangaoming@chinarun.com", "123456", 1, []user.RoleType{user.TechGuy})
	err = createUser("吴克兵", "wukebing@chinarun.com", "123456", 1, []user.RoleType{user.TechGuy})
	err = createUser("许航", "xuhang@chinarun.com", "123456", 1, []user.RoleType{user.TechGuy})
	err = createUser("张丙振", "zhangbingzhen@chinarun.com", "123456", 1, []user.RoleType{user.TechGuy})

	//CQ
	err = createUser("贾军伟", "runffjw@runff.com", "123456", 3, []user.RoleType{user.TechGuy})
	err = createUser("淦登杰", "gdj@runff.com", "123456", 3, []user.RoleType{user.TechGuy})
	err = createUser("李洁", "td@runff.com", "123456", 3, []user.RoleType{user.TechGuy})
	err = createUser("陈谷浩", "cgh@runff.com", "123456", 3, []user.RoleType{user.TechGuy})
	err = createUser("范文政", "fwz@runff.com", "123456", 3, []user.RoleType{user.TechGuy})
	err = createUser("汤程陈", "tcc@runff.com", "123456", 3, []user.RoleType{user.TechGuy})
	err = createUser("张永刚", "zyg@runff.com", "123456", 3, []user.RoleType{user.TechGuy})

	//制作人员--美术
	//SH
	err = createUser("孟祥太", "mengxiangtai@chinarun.com", "123456", 1, []user.RoleType{user.ArtGuy})
	err = createUser("贾考祥", "jiakaoxiang@chinarun.com", "123456", 1, []user.RoleType{user.ArtGuy})
	err = createUser("吴晨曦", "wuchenxi@chinarun.com", "123456", 1, []user.RoleType{user.ArtGuy})

	//YZ
	err = createUser("佘梦月", "chinarunsmy@chinarun.com", "123456", 2, []user.RoleType{user.ArtGuy})

	//fix
	err = createUser("徐婷婷", "chinarunxtt@chinarun.com", "123456", 2, []user.RoleType{user.BussinessMen})
	err = createUser("李易阳", "lyy@runff.com", "123456", 3, []user.RoleType{user.TechGuy})
	return
}
