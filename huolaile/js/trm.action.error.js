function trmErr(){
	this._string_ = new Array();
	var self = this;
}
/**
 * 判断service logic, SL
 * @param  data 错误时 callback返回的数据
 * @param dom_type string  value : list | obj | input | select | table 请求数据填充对应的dom元素的形式
 * @param dom string  请求数据填充对应的dom元素
 * @param img_name string 错误提示里面的图片
 * @param msg_dom string  错误提示对应的dom
 * @param msg string 错误消息-文本信息
 * @param page_dom 分页dom
 */
trmErr.prototype.SLNull = function(data,dom_type,dom,img_name,msg_dom,msg,page_dom) {
	
	if(dom!=null)
	dom.not(dom[0]).remove();
	
	if(!data || !data.reason || data.reason.reasoncode != TrmVar.GC('NOEXIST'))
		return;
	
	var dom_html='<div class="error-content"><div class="'+img_name+'"></div>'+
	             '<h5>'+TrmVar.GC('MSG_BOW_NULL_SEARCH')+msg+'</h5>'+
	             '</div>';
	
	var select_html = "<option style='color:#999'>还没有帮您查到下拉信息...</option>";
	
	
	switch(dom_type){
		case "list":
			if(page_dom!=null) page_dom.remove();
			
			dom.empty()
		       .append(dom_html);      					
			break;
			
		case "table":			
			dom.empty();
		    msg_dom.html(dom_html);
		    msg_dom.next(".bottom-bar").hide();
		    if(page_dom!=null)
		    page_dom.pagination('updateItems', 1);
			break;		    		
			
		case "obj":
			if(dom!=null)
			dom.html(dom_html);  
			
			Confirm.show('提示',"warning", "未找到指定"+msg, 
					      {'关闭': {'callback': function() {Confirm.hide();}}}
			            );
			break;
			
		case "input":		
			break;
			
		case "select":
			dom.html(select_html);
			break;
		default: 
			break;
	}
}

/**
 * 数据操作错误处理 data manipulation, DM -base
 * @param  data 错误时 callback返回的数据
 * @param dom_type string  value : list | obj | input | select | table 请求数据填充对应的dom元素的形式
 * 
 */
trmErr.prototype.DMBase = function(data,dom_type) {
	if(!data || !data.reason || data.reason.reasoncode != TrmVar.GC('STATUE_EXCEP'))
		return;
	Confirm.show('提示',"warning", '操作失败，请重试!', {
		'确定': {'callback': function() {Confirm.hide();}}
    });
}

/**
 * 数据提交错误处理 data submission, DS -base
 * @param  data 错误时 callback返回的数据
 */
trmErr.prototype.DSBase = function(data){
	if(!data || !data.reason || data.reason.reasoncode != TrmVar.GC('STATUE_EXCEP'))
		return;
	Confirm.show('提示',"warning", '操作失败，请重试!', {
		'确定': {'callback': function() {Confirm.hide();}}
    });
}

/**
 * 数据提交错误处理 data submission, DS -Verify
 * @param  data 错误时 callback返回的数据
 */
trmErr.prototype.DSVerify = function(data){
	if(!data || !data.reason || data.errorType != "1")
		return;
	
	if(typeof(data.filedlist) != "undefined" && data.filedlist.length>0){
		var str = "";
		$.each(data.filedlist, function(index, item){
			if(typeof(item.message) != "undefined" && item.message != null){
				str = '<label class="error"><span class="valid_message_select"></span>'+item.message+'</label>';
			}else{
				str = '<label class="error"><span class="valid_message_select"></span>输入有误，请检查</label>';
			}
			
			if(typeof(item.fieldName) != "undefined" && item.fieldName != null){
				$("#"+item.fieldName+"_filed").parent().find(".error-tips").empty();
				$("#"+item.fieldName+"_filed").parent().find(".error-tips").html(str);
			}else if(typeof(item.message) != "undefined" && item.message != null){
				Confirm.show('提示',"warning", item.message, {
					'关闭': {
						'callback': function() {
							Confirm.hide();
						}
					}
				});
			}
		});
	}else{
		Confirm.show('提示',"warning", "您输入的数据有误，请检查！", {
			'关闭': {'callback': function() {Confirm.hide();}
			}
		});
	}

}

/**
 * 重复提交数据处理 data submission, DS -Repeat
 * @param  data 错误时 callback返回的数据
 * @param  dom_type
 * @param  msg
 */
trmErr.prototype.DSRepeat = function(data,dom_type,msg){
	
	if(!data || !data.reason || data.reason.reasoncode != TrmVar.GC('REPEAT_ADD'))
		return;
	switch(dom_type){	   
	    case "obj":
			Confirm.show('提示',"warning", "您已"+msg+"!", {
				         '关闭':{'callback': function() {Confirm.hide();}}
			});
		    break;
		    
		default:
			break;
	}
}

/**
 * 重复提交数据处理 data submission, DS - NoLogin
 * @param  data 错误时 callback返回的数据
 * @param  dom_type
 * @param  msg
 * 后期可以考虑放入 action.ajax中
 */
trmErr.prototype.DSNoLogin = function(data,dom_type,msg){
	
	if(!data || !data.reason || data.reason.reasoncode != TrmVar.GC('NOT_LOGIN'))
		return;
	switch(dom_type){	   
		case "obj":
			Confirm.show('提示',"warning", "请您登录后再执行"+msg+"操作!", {
				         '关闭':{'callback': function() {Confirm.hide();}}
			});
	        break;
	    
		default:
			break;
	}
}



