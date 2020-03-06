var TrmVar = (function() {  
	
  // Private static attributes.
  var constants = {
		  
	ERR_TYPE_BASE : "0",    //基础业务错误
	ERR_TYPE_VERIFY : "1",  //校验错误
	ERR_TYPE_ACCESSRIGHT : "2",   //访问权限
	ERR_TYPE_C : "4",   //C数据
		  
    BIZ : "BUSI_ERR_9999",  //业务级别
    SYS : "016",  //系统级别
    OTHER : "OTHER_SYS_ERR",  //系统级别
    RUNERR : "JAVA_RUNTIME_ERR",  //系统级别
    
    NOEXIST : "2048",  //查询数据不存在
    DATE_ILLEGAL: "2068", //数据校验不合法 (目前特指时间校验)
    STATUE_EXCEP: "2088", //状态转换异常
    REPEAT_ADD: "3056", //重复添加
    NOT_LOGIN: "4022", //用户未登陆    
    TRADE_EXCEP: "4241", //交易类型
    TRADE_NOT_COMPLETE: "6248", //订单支付未完成
    
    UC_NOT_LOGIN: "3001", //前端用户未登录
    MGT_NOT_LOGIN: "3002", //后端用户未登录
    NOT_REAL_CERT: "3003", //未实名认证
    
    C_EMAIL_SENDFAIL : "E10000", //邮件发送失败    
    C_VERIFCODE_PASSFAIL : "E10112", //验证码错误或已失效
    C_ORIGINALPSW_FAIL : "E10111", //原密码输入错误
    C_PARA_NULL : "E11000", //传入参数不能为空
    C_EMAIL_NOTSAME : "E10113", //输入邮箱和注册邮箱不一致
    C_PSW_NOTSAME : "E10114", //两次密码不一致
    C_ACCPSW_ERROR : "E10115", //账号不存在或密码错误
    C_Mail_BEEN_REG : "E10116", //该邮箱已被注册
    C_ACC_BEEN_REG : "E10117", //该账号已被注册
    
    C_REDIS_RWFAIL : "E10110", //Redis读写信息失败
    C_NOT_HASRIGHT : "E10010", //用户没有权限  
    MSG_LOG: "系统其他错误类型",
    MSG_LOG_SYS: "系统不可用，或请求的数据格式有误",
    
    MSG_LOG_NOEXIST : "查询数据不存在",  //查询数据不存在
    MSG_LOG_DATE_ILLEGAL: "数据校验不合法",
    MSG_LOG_STATUE_EXCEP: "状态转换异常",
    MSG_LOG_NOT_LOGIN: "用户未登陆",
    MSG_LOG_TRADE_EXCEP: "交易中订单异常操作",
    
    MSG_BOW_NULL_SEARCH: "抱歉，没有找到相关",
    MSG_BOW_NULL_PERSON: "您还没有",
    MSG_BOW_NULL_BASE: "暂无",
    
    MSG_C_LOG : "C-接口错误",
    MSG_C_LOG_REDIS_RWFAIL : "C-Redis读写信息失败", 
    MSG_C_LOG_NOT_HASRIGHT : "C-用户没有权限 "
    
    
  }
  var VarObj={};
  
  // Private static method.
  VarObj.GC=function(name){

    return constants[name];
  }
  return VarObj
})();