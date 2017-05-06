//==================================================================
//创建时间：2017-4-23 首次创建
//功能描述：状态码
//创建人：郭世江
//修改记录：若要修改请记录
//==================================================================
package types
const (
	UpdateCheckStatus4DepartmentsErr      = 1000        //更新状态失败
	UpdateCheckupResultErr                 = 1001        //更新结果失败
	WriteLogErr                             = 1002        //写日志失败
	RemoveFileErr                           = 1003        //删除文件失败
	GetAccountInfoErr                      = 1004        //获取操作员信息失败
	GetQueryCodeErr                        = 1005        //获取查询码失败
	InsertDataDBErr                        = 1007        //将数据插入数据库中失败
	XmlHearDataErr                         = 1008
)

const (
	UpdateCheckStatus4DepartmentsSucc      = 2000       //更新状态成功
	UpdateCheckupResultErrSucc             = 2001        //更新结果成功
	WriteLogSucc                             = 2002        //写日志成功
	RemoveFileSucc                           = 2003        //删除文件成功
	GetAccountInfoSucc                      = 2004        //获取操作员信息成功
	GetQueryCodeSucc                        = 2005        //获取查询码成功
	InsertDataDBSucc                        = 2007        //将数据插入数据库中成功
	XmlHearDataSucc                         = 2008       //解析XML失败
)