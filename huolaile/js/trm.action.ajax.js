/**
 * @version: 1.0.1
 * @author:
 * @date: 2016-06-19
 * @function: ajax request 封装
 * @dependence: jquery-2.2.3.min.js
 *
 */

/**
 * @function ajaxCallScript
 * @param  option.url : 请求的url
 *          option.requestData : 发送请求的参数
 *          option.hasload : 是否使用loading状态 默认为false
 *          option.loadDiv : 显示loading状态的Div
 *            option.timeout : 超时设置，默认为10000
 *          option.successHandler : 请求成功的回调函数
 *          option.timeoutHandle : 请求超时的回调函数
 */
var ajaxCallScript = function (option) {

    this.url = option.url;

    this.requestData = option.requestData;

    if (typeof(this.requestData) != "undefined")
        this.requestData.ajaxIndicator = 'yes';
    else
        this.requestData = {"ajaxIndicator": "yes"};

    this.hasload = false || option.hasload;

    this.loadDiv = option.loadDiv;

    option.timeout > 300000 ? this.timeout = option.timeout
        : this.timeout = 300000;

    if (localStorage.getItem('token'))
        requestData.token = localStorage.getItem('token');

    var hasErrHandler = typeof(option.errorHandler) != "undefined";

    var jqXHR = jqXHRFactory();

    jqXHR.done(function (data) {
        //success
        if (data.hasOwnProperty("responsecode")) {
            if (data.responsecode == TrmVar.GC('BIZ')) {
                switch (data.errorType) {

                    case TrmVar.GC('ERR_TYPE_BASE'):
                        BIZ_BASE_ErrorHandle(data.reason);
                        hasErrHandler ? errorHandlerProxy(option.errorHandler)(data) : '';
                        break;

                    case TrmVar.GC('ERR_TYPE_VERIFY'):
                        console.info("数据校验有错");
                        hasErrHandler ? errorHandlerProxy(option.errorHandler)(data) : '';
                        break;

                    case TrmVar.GC('ERR_TYPE_ACCESSRIGHT'):
                        console.info("请求权限有误");
                        RIGHT_BASE_ErrorHandle(data.reason);
                        break;

                    case TrmVar.GC('ERR_TYPE_C'):
                        BIZ_C_ErrorHandle(data.reason);
                        hasErrHandler ? errorHandlerProxy(option.errorHandler)(data) : '';
                        break;

                    default:
                        break;

                }

            } else if (data.responsecode == TrmVar.GC('SYS')) {
                console.info(TrmVar.GC('MSG_LOG_SYS'));
                //建议处理系统不可用的页面
                //window.location.href="/busy.html";
            } else if (data.responsecode == TrmVar.GC('OTHER')) {
                console.info(TrmVar.GC('MSG_LOG'));
                hasErrHandler ? errorHandlerProxy(option.errorHandler)(data) : '';
            } else if (data.responsecode == TrmVar.GC('RUNERR')) {
                console.info(TrmVar.GC('MSG_LOG'));
                hasErrHandler ? errorHandlerProxy(option.errorHandler)(data) : '';
            } else {
                console.info(TrmVar.GC('MSG_LOG'));
                //window.location.href="/busy.html";
            }

        } else {
            successHandlerProxy(option.successHandler)(data);
        }

    }).fail(function (jqxhr, textStatus, error) {
        //fail
        var err = textStatus + ',' + error;
        console.warn("Request Fail:" + err);
        // console.error("Request Fail:" + error);
        //  if(textStatus=="error") window.location.href="/nofind.html";

    }).always(function (jqxhr, textStatus, error) {
        //complete
    });

    return jqXHR;

    /**
     *
     * @param errorHandler
     * @returns {Function}
     */
    function errorHandlerProxy(errorHandler) {
        if (typeof(errorHandler) != "undefined")
            return function (data) {
                errorHandler(data);
            }
    }

    /**
     *
     * @param successHandler
     * @returns {Function}
     */
    function successHandlerProxy(successHandler) {
        return function (data) {
            successHandler(data);
        }
    }

    /**
     *
     * @param successHandler
     * @returns {Function}
     */
    function timeoutHandlerProxy(timeoutHandler) {
        return function () {
            timeoutHandler();
        }
    }

    /**
     * @returns {Function}
     */
    function jqXHRFactory() {
        var hasload = this.hasload;
        var loadDiv = this.loadDiv;
        return $.ajax({
            url: this.url,
            type: "post",
            data: JSON.stringify(this.requestData),
            dataType: this.dataType || "json",
            contentType: 'application/json; charset=UTF-8',
            timeout: this.timeout,
            beforeSend: function () {

//        	   var reLA='<div class="loading-content"><div class="la-square-jelly-box la-2x"><div></div><div></div></div><h6>正在加载数据...</h6></div>';
//               
//        	   if (hasload){
//        		   loadDiv.before(reLA);
//        		   loadDiv.fadeIn();
//        	   }

            },
            complete: function () {
//               if (hasload)
//            	   loadDiv.prev(".loading-content").fadeOut(300);
//               loadDiv.fadeIn();
            }
        });
    }

    /**
     * A-业务逻辑处理
     *
     */
    function BIZ_BASE_ErrorHandle(reason) {
        var log_info = "";

        switch (reason.reasoncode) {

            case TrmVar.GC('NOEXIST'):
                log_info += TrmVar.GC('MSG_LOG_NOEXIST');
                reason.desc != null ? log_info += "," + reason.desc : '';
                console.info(log_info);
                break;
            case TrmVar.GC('DATE_ILLEGAL'):
                log_info += TrmVar.GC('MSG_LOG_DATE_ILLEGAL');
                reason.desc != null ? log_info += "," + reason.desc : '';
                console.info(log_info);
                break;
            case TrmVar.GC('STATUE_EXCEP'):
                log_info += TrmVar.GC('MSG_LOG_STATUE_EXCEP');
                reason.desc != null ? log_info += "," + reason.desc : '';
                console.info(log_info);
                break;
            case TrmVar.GC('REPEAT_ADD'):
                log_info += TrmVar.GC('MSG_LOG_REPEAT_ADD');
                reason.desc != null ? log_info += "," + reason.desc : '';
                console.info(log_info);
                break;
            case TrmVar.GC('TRADE_EXCEP'):
                log_info += TrmVar.GC('MSG_LOG_TRADE_EXCEP');
                reason.desc != null ? log_info += "," + reason.desc : '';
                console.info(log_info);
                break;
            case TrmVar.GC('NOT_LOGIN'):
                log_info += TrmVar.GC('MSG_LOG_NOT_LOGIN');
                reason.desc != null ? log_info += "," + reason.desc : '';
                console.info(log_info);
                break;
            default:
                console.info(TrmVar.GC('MSG_LOG_BIZ'));
                break;

        }

    }

    /**
     * C-业务逻辑处理
     *
     */
    function BIZ_C_ErrorHandle(reason) {
        var log_info = "";

        switch (reason.reasoncode) {

            case TrmVar.GC('C_REDIS_RWFAIL'):
                log_info += TrmVar.GC('MSG_C_LOG_REDIS_RWFAIL');
                reason.desc != null ? log_info += "," + reason.desc : '';
                console.info(log_info);
                break;
            case TrmVar.GC('C_NOT_HASRIGHT'):
                log_info += TrmVar.GC('MSG_C_LOG_NOT_HASRIGHT');
                reason.desc != null ? log_info += "," + reason.desc : '';
                console.info(log_info);
                break;

            default:
                console.info(TrmVar.GC('MSG_C_LOG'));
                break;

        }

    }

    function RIGHT_BASE_ErrorHandle(reason) {
        switch (reason.reasoncode) {
            case TrmVar.GC('UC_NOT_LOGIN'):
                window.location.href = "/views/content/login/front/login.jsp";
                break;
            case TrmVar.GC('NOT_REAL_CERT'):
                window.location.href = "/views/content/apply/front/apply.jsp";
                break;
            default:
                break;
        }
    }

}
