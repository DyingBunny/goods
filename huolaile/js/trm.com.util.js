function trmUtil(){
	this._string_ = new Array();
	var self = this;
}
/*
 * 多个字符串拼接为数组
 */
trmUtil.prototype.Append = function(str) {
	this._string_.push(str);
}

/*
 * 数组转化为字符串
 */
trmUtil.prototype.toString = function() {
	return this._string_.join("");
}

/*
 * 把一个字符串变成链接
 */
trmUtil.prototype.Filter = function(str) {

	var urlReg = /http(s)?:\/\/([\w-]+\.)+[\w-]+(\/[\w- .\/?%&=])?[^ <>\[\]*(){}\u4E00-\u9FA5]+/gi; //lio 2012-4-25 eidt   //         /^[\u4e00-\u9fa5\w]+$/;\u4E00-\u9FA5

	return str.replace(urlReg, function(m) {
		return '<a target="_blank" href="' + m + '">' + m + '</a>';
	});

}

/*
 * 验证一个字符串是否是汉字
 */
trmUtil.prototype.isZHCN = function(str) {
	var p = /^[\u4e00-\u9fa5\w]+$/;
	return p.test(str);
}

/*
 * 验证一个字符串是否是数字
 */
trmUtil.prototype.isNum = function(str) {
	var p = /^\d+(\.\d+)?$/;

	return p.test(str);
}
trmUtil.prototype.cutStr = function(str,len){
	var str1 = str;
	str1 = str1.replace(/(\n)/g, "");  
	str1 = str1.replace(/(\t)/g, "");  
	str1 = str1.replace(/(\r)/g, "");  
	str1 = str1.replace(/<\/?[^>]*>/g, "");  
	//str1 = str1.replace(/\s*/g, ""); 
	str1 = str1.cutBytes(len,'……');
	
	var ser_arr1,ser_arr2;
	ser_arr1=str.match(/<span style='background:#fed10a;'>.*?<\/span>/gm);
	if(ser_arr1 == null) return str1;
	
	ser_arr2=ser_arr1.slice(0);
	for(var i=0;i<ser_arr2.length;i++){ 
		ser_arr2[i]=ser_arr2[i].replace(/<\/?[^>]*>/g, '');
	};

	var k=0,b_i=0;
	for(var j=0; j<ser_arr2.length; j++) { 
		b_i=str1.indexOf(ser_arr2[j],k);
		if(b_i!=-1){
			str1=str1.substring(0,b_i)+ser_arr1[j]+str1.substring(b_i+ser_arr2[j].length,str1.length);
		    k=b_i+ser_arr1[j].length;
		}
	};
	return str1;
}
/*
 * 验证一个字符串是否是纯英文
 */
trmUtil.prototype.isEnglish = function(str) {
	var p = /^[a-zA-Z., ]+$/;
	return p.test(str);
}

/*
 * 判断是否为对象类型
 */
trmUtil.prototype.isObject = function(obj) {
	return (typeof obj == 'object') && obj.constructor == Object;
}

/*
 * 验证字符串是否不包含特殊字符 返回bool
 */
trmUtil.prototype.isUnSymbols = function(str) {
	var p = /^[\u4e00-\u9fa5\w \.,()，ê?。¡ê（ê¡§）ê?]+$/;
	return p.test(str);
}

/* 
说明：对复选框的全选或不选 
参数state:输入值 1 全选 2 全不选 
参数name:限定全选范围为指定名称的checkbox
*/
trmUtil.prototype.CheckBoxAllChange = function(state, name) {
	try {
		var checks = document.getElementsByTagName(name);
		var i = 0;
		var length = checks.length;
		var flag = true;
		if (state == 1) {
			flag = true;
		}
		if (state == 0) {
			flag = false;
		}
		for (i; i < length; i++) {
			if (checks[i].type == "checkbox") {
				checks[i].checked = flag;
			}
		}
	} catch (e) {
		window.alert(e.message);
	}
}

/*
 * 获取URL的详细信息
 */
trmUtil.prototype.parseURL = function(url) {
	var a = document.createElement('a');
	a.href = url;
	return {
		source: url,
		protocol: a.protocol.replace(':', ''),
		host: a.hostname,
		port: a.port,
		query: a.search,
		params: (function() {
			var ret = {},
				seg = a.search.replace(/^\?/, '').split('&'),
				len = seg.length,
				i = 0,
				s;
			for (; i < len; i++) {
				if (!seg[i]) {
					continue;
				}
				s = seg[i].split('=');
				ret[s[0]] = decodeURI(s[1]);
			}
			return ret;
		})(),
		file: (a.pathname.match(/\/([^\/?#]+)$/i) || [, ''])[1],
		hash: a.hash.replace('#', ''),
		path: a.pathname.replace(/^([^\/])/, '/$1'),
		relative: (a.href.match(/tps?:\/\/[^\/]+(.+)/) || [, ''])[1],
		segments: a.pathname.replace(/^\//, '').split('/')
	};
}

trmUtil.prototype.removeParameter = function(destiny, par) {
	// var reg=new RegExp("\\\? | &"+par+"= ([^&]+)(&|&)","i");
	// return destiny.replace(reg,"");
	var str = "";

	if (destiny.indexOf('?') != -1)
		str = destiny.substr(destiny.indexOf('?') + 1);
	else
		return destiny;
	var arr = "";
	var returnurl = "";
	var setparam = "";
	if (str.indexOf('&') != -1) {
		arr = str.split('&');
		//for(i in arr) {
		for (var i = 0; i < arr.length; i++){
			if (arr[i].split('=')[0] != par) {
				returnurl = returnurl + arr[i].split('=')[0] + "="
						+ arr[i].split('=')[1] + "&";
			}
		}
		return destiny.substr(0, destiny.indexOf('?')) + "?"
				+ returnurl.substr(0, returnurl.length - 1);
	} else {
		arr = str.split('=');
		if (arr[0] == par)
			return destiny.substr(0, destiny.indexOf('?'));
		else
			return destiny;
	}
}

trmUtil.prototype.changeURLPar = function(destiny, par, par_value) {
	// var destiny = window.location.href.toString();
	var pattern = par + '=([^&]*)';
	var replaceText = par + '=' + par_value;
	if(destiny.match(pattern)){
		var re=eval('/('+ par+'=)([^&]*)/gi');
	    var tmp = destiny.replace(re,par+'='+par_value);
	    
//		var tmp = '/\\' + par + '=[^&]*/';
//		tmp = destiny.replace(eval(tmp), replaceText);
	
		return tmp;
	}else{
		if(destiny.match('[\?]')){
			return destiny + '&' + replaceText;
		}else{
			return destiny + '?' + replaceText;
		}
	}
	destiny +=  '\n' + par + '\n' + par_value;
    return destiny;
} 

/*
 * 格式化电话号码  例如  ： 133 020 2222
 */
trmUtil.prototype.formatTel = function(num) {
	var reg = /^(\d{3})(\d{4})(\d{4})$/;
	var matches = reg.exec(num);
	if(!matches)return num;
	return  matches[1] + ' ' + matches[2] + ' ' + matches[3];
}


/*
 * 取得文件扩展名
 */
trmUtil.prototype.getFileExt = function(path) {
	var tmp = path;
	tmp = tmp.substring(tmp.lastIndexOf(".") + 1);
	return tmp.toUpperCase();
}

/*
 * 检查是否为金额字符串
 */
trmUtil.prototype.isMoneyStr = function(str) {
	var self = this;
	if (str.Trim().isNull())
		return false;

	var str1;
	var dotnum = 0;
	//进行检测
	for (i = 0; i < str.length; i++) {
		str1 = str.substr(i, 1);
		//检测是否为负
		if (i == 0) {
			if ((str1 != "-") && (!self.isNum(str1)))
				return false;
		} else {
			if ((str1 != ".") && (!self.isNum(str1)))
				return false;
			if (str1 == ".")
				dotnum++;
			if (dotnum > 1)
				return false;
		}
	}
	return true;
}

/* 
 * formatMoney(s,type) 
 * 功能：金额按千位逗号分割 
 * 参数：s，需要格式化的金额数值. 
 * 参数：type,判断格式化后的金额是否需要小数位. 
 * 返回：返回格式化后的数值字符串. 
 */
trmUtil.prototype.formatMoney = function(s, type) {
	if (/[^0-9\.]/.test(s))
		return "0";
	if (s == null || s == "")
		return "0";
	s = s.toString().replace(/^(\d*)$/, "$1.");
	s = (s + "00").replace(/(\d*\.\d\d)\d*/, "$1");
//	s = s.replace(".", ",");	/*wangxun modify 2016/08/11 */
//	var re = /(\d)(\d{3},)/;
//	while (re.test(s))
//		s = s.replace(re, "$1,$2");
//	s = s.replace(/,(\d\d)$/, ".$1");
	if (type == 0) { // 不带小数位(默认是有小数位)  
		var a = s.split(".");
		if (a[1] == "00") {
			s = a[0];
		}
	}
	return s;
}

/* 
 * rmoney(s) 
 * 功能：金额还原为普通数字
 * 参数：s，需要还原的金额数值. 
 * 返回：返回还原后的数值字符串. 
 */
trmUtil.prototype.rmoney = function(s) {
	return parseFloat(s.replace(/[^\d\.-]/g, ""));
}


trmUtil.prototype.formatFloat = function(f, digit) {
    var m = Math.pow(10, digit);
    return parseInt(f * m, 10) / m;
} 

/**
 *判断str1字符串是否以str2为结束
 *@param str1 原字符串
 *@param str2 子串
 *@return 是-true,否-false
 */
trmUtil.prototype.endWith = function(str1, str2) {
	if (str1 == null || str2 == null) {
		return false;
	}
	if (str1.length < str2.length) {
		return false;
	} else if (str1 == str2) {
		return true;
	} else if (str1.substring(str1.length - str2.length) == str2) {
		return true;
	}
	return false;
}

/**
 *判断str1字符串是否以str2为开头
 *@param str1 原字符串
 *@param str2 子串
 *@return 是-true,否-false
 */
trmUtil.prototype.startWith = function(str1, str2) {
	if (str1 == null || str2 == null) {
		return false;
	}
	if (str1.length < str2.length) {
		return false;
	} else if (str1 == str2) {
		return true;
	} else if (str1.substr(0, str2.length) == str2) {
		return true;
	}
	return false;
}

/*
 * 保留字符串中的数字
 */
trmUtil.prototype.getNum = function(text) {
	return parseInt(text.replace(/[^\d.]/g, ''));
}

/*
 * 获取当前浏览器名称及版本信息
 */
trmUtil.prototype.getBrowserInfo = function() {
	var agent = navigator.userAgent.toLowerCase();

	var regStr_ie = /msie [\d.]+;/gi;
	var regStr_ff = /firefox\/[\d.]+/gi
	var regStr_chrome = /chrome\/[\d.]+/gi;
	var regStr_saf = /safari\/[\d.]+/gi;
	//IE
	if (agent.indexOf("msie") > 0) {
		return {
			"name": agent.match(regStr_ie).toString().split(" ")[0],
			"version": agent.match(regStr_ie).toString().split(" ")[1]
		}
	}

	//firefox
	if (agent.indexOf("firefox") > 0) {
		return {
			"name": agent.match(regStr_ff).toString().split("/")[0],
			"version": agent.match(regStr_ff).toString().split("/")[1]
		}
	}

	//Chrome
	if (agent.indexOf("chrome") > 0) {
		return {
			"name": agent.match(regStr_chrome).toString().split("/")[0],
			"version": agent.match(regStr_chrome).toString().split("/")[1]
		}
	}

	//Safari
	if (agent.indexOf("safari") > 0 && agent.indexOf("chrome") < 0) {
		return {
			"name": agent.match(regStr_saf).toString().split("/")[0],
			"version": agent.match(regStr_saf).toString().split("/")[1]
		}
	}
}

/**
 *返回筛选列表选中的参数
 *@param activeClass 选中的参数class
 *@return 根据activeClass筛选的参数列表，JSON格式
 */
trmUtil.prototype.getActiveParams = function(activeClass) {
	if (!activeClass) return false;
	var obj = window.document.getElementsByClassName(activeClass);
	var params = {};
	for (var i = 0; i < obj.length; i++) {
		var id = obj[i].id;
		if (!params[obj[i].parentElement.id]) {
			params[obj[i].parentElement.id] = {};
		}
		params[obj[i].parentElement.id][id] = id;
	}
	return params;
}

/*
 * 解决IE8之类不支持getElementsByClassName
 */
if (!document.getElementsByClassName) {
	document.getElementsByClassName = function(className, element) {
		var children = (element || document).getElementsByTagName('*');
		var elements = new Array();
		for (var i = 0; i < children.length; i++) {
			var child = children[i];
			var classNames = child.className.split(' ');
			for (var j = 0; j < classNames.length; j++) {
				if (classNames[j] == className) {
					elements.push(child);
					break;
				}
			}
		}
		return elements;
	};
}

/**
 *按照指定参数对DOM对象进行排序
 *@param obj 排序对象
 *@param sortClass 排序参数
 *@return 排序后的DOM对象
 */
trmUtil.prototype.sortDOM = function(obj, sortClass, reverse) {
	var self = this;
	if (!obj || !sortClass) return obj;
	var param = [];
	for (var i = 0; i < obj.length; i++) {
		param[i] = obj[i];
	}
	return param.sort(self.compare(sortClass, reverse));
}

trmUtil.prototype.compare = function(propertyName, reverse) {
	var self = this;
	if (reverse) {
		return function(object1, object2) {
			var a = new DateUtil();
			var value1, value2;
			if (self.isNum(object1.getAttribute(propertyName))) {
				value1 = Number(object1.getAttribute(propertyName));
				value2 = Number(object2.getAttribute(propertyName));
			} else if (a.isDate(object1.getAttribute(propertyName), "yyyy-MM-dd")) {
				value1 = a.dateToLong(a.strToDate(object1.getAttribute(propertyName)));
				value2 = a.dateToLong(a.strToDate(object2.getAttribute(propertyName)));
			}

			if (value2 < value1) {
				return -1;
			} else if (value2 > value1) {
				return 1;
			} else {
				return 0;
			}
		}
	} else {
		return function(object1, object2) {
			var a = new DateUtil();
			var value1, value2;
			if (self.isNum(object1.getAttribute(propertyName))) {
				value1 = Number(object1.getAttribute(propertyName));
				value2 = Number(object2.getAttribute(propertyName));
			} else if (a.isDate(object1.getAttribute(propertyName), "yyyy-MM-dd")) {
				value1 = a.dateToLong(a.strToDate(object1.getAttribute(propertyName)));
				value2 = a.dateToLong(a.strToDate(object2.getAttribute(propertyName)));
			}
			if (value2 > value1) {
				return -1;
			} else if (value2 < value1) {
				return 1;
			} else {
				return 0;
			}
		}
	}
}

/*
 * 检测屏幕分辨率（高度大于768px为true）
 */
trmUtil.prototype.checkScreenSize = function() {
	var browser = navigator.appName;
	if (browser != "Microsoft Internet Explorer") return (window.innerHeight > 768) ? true : false;
	var b_version = navigator.appVersion;
	var version = b_version.split(";");
	var trim_Version = version[1].replace(/[ ]/g, "");
	if (browser == "Microsoft Internet Explorer" && (trim_Version == "MSIE8.0" || trim_Version == "MSIE7.0" || trim_Version == "MSIE6.0")) {
		return (screen.availHeight > 768) ? true : false;
	} else {
		return (window.innerHeight > 768) ? true : false;
	}
//	// 获取窗口宽度
//	if (window.innerWidth)
//		winWidth = window.innerWidth;
//	else if ((document.body) && (document.body.clientWidth))
//		winWidth = document.body.clientWidth;
//	// 获取窗口高度
//	if (window.innerHeight)
//		winHeight = window.innerHeight;
//	else if ((document.body) && (document.body.clientHeight))
//		winHeight = document.body.clientHeight;
//	// 通过深入 Document 内部对 body 进行检测，获取窗口大小
//	if (document.documentElement && document.documentElement.clientHeight && document.documentElement.clientWidth) {
//		winHeight = document.documentElement.clientHeight;
//		winWidth = document.documentElement.clientWidth;
//	}
}

/*
 * 获取字符串长度 区分中英文, 中文两个字节
 */
String.prototype.getBytes = function() {
	var cArr = this.match(/[^\x00-\xff]/ig);
	return this.length + (cArr == null ? 0 : cArr.length);
}

/*
 * 截取字符串长度 区分中英文, 中文两个字节. 超出部分中指定字符串代替 需引用 String.prototype.getBytes
 */
String.prototype.cutBytes = function(strLen, replaceStr) {
	var str = this.toString();
	if (str.length <= strLen)
		return str;
	var returnStr = "";
	var tempLen = 0;
	for (var i = 0; i < str.length; i++) {
		returnStr += str[i];
		tempLen++;
		if (tempLen >= strLen) {
			return i + 1 < str.length ? returnStr + replaceStr : returnStr;
		}
	}
	return "";
};

/*
 * trim去掉字符串两边的指定字符,默去空格
 */
String.prototype.Trim = function(str) {
	if (!str) {
		str = '\\s';
	} else {
		if (str == '\\') {
			str = '\\\\';
		} else if (str == ',' || str == '|' || str == ';') {
			str = '\\' + str;
		} else {
			str = '\\s';
		}
	}
	eval('var reg=/(^' + str + '+)|(' + str + '+$)/g;');
	return this.replace(reg, '');
};

String.prototype.trim = function(str) {
	return this.Trim(str);
};

/*
 * 判断一个字符串是否为NULL或者空字符串
 */
String.prototype.isNull = function() {
	return this == null || this.trim().length == 0;
}

/*
 * 判断字符串是否相等
 */
String.prototype.equals = function(str) {
	return this == str;
}

/*
 * 忽略大小写判断字符串是否相同
 */
String.prototype.isEqualsIgnorecase = function(str) {
	if (this.toUpperCase() == str.toUpperCase()) {
		return true;
	}
	return false;
}

/*
 * 字符串截取后面加入...
 */
String.prototype.interceptString = function(len) {
	a
	if (this.length > len) {
		return this.substring(0, length - 1) + "...";
	} else {
		return this;
	}
}

/*
 * 根据数据取得在数组中的索引
 */
Array.prototype.getIndex = function(obj) {
	var objStr = obj.toString();
	for (var i = 0; i < this.length; i++) {
		if (objStr == this[i] || objStr.equals(this[i])) {
			return i;
		}
	}
	return -1;
}

/*
 * 移除数组中的某元素
 */
Array.prototype.remove = function(obj) {
	var objStr = obj.toString();
	for (var i = 0; i < this.length; i++) {
		if (objStr.equals(this[i])) {
			this.splice(i, 1);
			break;
		}
	}
	return this;
};

/*
 * 判断元素是否在数组中
 */
Array.prototype.contains = function(obj) {
	var objStr = obj.toString();
	for (var i = 0; i < this.length; i++) {
		if (objStr == this[i] || objStr.equals(this[i])) {
			return true;
		}
	}
	return false;
};

/*
 * 日期处理工具类
 */
var DateUtil = function() {
	/**
	 *	变量是否为数字
	 */
	this.isNumber = function(str) {
		var regExp = /^\d+$/g;
		return regExp.test(str);
	}

	/**
	 * 判断闰年
	 * @param date Date日期对象
	 * @return boolean true 或false
	 */
	this.isLeapYear = function(date) {
		return (0 == date.getYear() % 4 && ((date.getYear() % 100 != 0) || (date.getYear() % 400 == 0)));
	}

	/**
	 * 日期对象转换为指定格式的字符串
	 * @param f 日期格式,格式定义如下 yyyy-MM-dd HH:mm:ss
	 * @param date Date日期对象, 如果缺省，则为当前时间
	 *
	 * YYYY/yyyy/YY/yy 表示年份  
	 * MM/M 月份  
	 * W/w 星期  
	 * dd/DD/d/D 日期  
	 * hh/HH/h/H 时间  
	 * mm/m 分钟  
	 * ss/SS/s/S 秒  
	 * @return string 指定格式的时间字符串
	 */
	this.dateToStr = function(formatStr, date) {
		formatStr = arguments[0] || "yyyy-MM-dd HH:mm:ss";
		date = arguments[1] || new Date();
		var str = formatStr;
		var Week = ['日', '一', '二', '三', '四', '五', '六'];
		str = str.replace(/yyyy|YYYY/, date.getFullYear());
		str = str.replace(/yy|YY/, (date.getYear() % 100) > 9 ? (date.getYear() % 100).toString() : '0' + (date.getYear() % 100));
		str = str.replace(/MM/, date.getMonth() > 9 ? (date.getMonth() + 1) : '0' + (date.getMonth() + 1));
		str = str.replace(/M/g, date.getMonth());
		str = str.replace(/w|W/g, Week[date.getDay()]);

		str = str.replace(/dd|DD/, date.getDate() > 9 ? date.getDate().toString() : '0' + date.getDate());
		str = str.replace(/d|D/g, date.getDate());

		str = str.replace(/hh|HH/, date.getHours() > 9 ? date.getHours().toString() : '0' + date.getHours());
		str = str.replace(/h|H/g, date.getHours());
		str = str.replace(/mm/, date.getMinutes() > 9 ? date.getMinutes().toString() : '0' + date.getMinutes());
		str = str.replace(/m/g, date.getMinutes());

		str = str.replace(/ss|SS/, date.getSeconds() > 9 ? date.getSeconds().toString() : '0' + date.getSeconds());
		str = str.replace(/s|S/g, date.getSeconds());

		return str;
	}

	/**
	 * 日期计算  
	 * @param strInterval string  可选值 y 年 m月 d日 w星期 ww周 h时 n分 s秒  
	 * @param num int
	 * @param date Date 日期对象
	 * @return Date 返回日期对象
	 */
	this.dateAdd = function(strInterval, num, date) {
		date = arguments[2] || new Date();
		switch (strInterval) {
			case 's':
				return new Date(date.getTime() + (1000 * num));
			case 'n':
				return new Date(date.getTime() + (60000 * num));
			case 'h':
				return new Date(date.getTime() + (3600000 * num));
			case 'd':
				return new Date(date.getTime() + (86400000 * num));
			case 'w':
				return new Date(date.getTime() + ((86400000 * 7) * num));
			case 'm':
				return new Date(date.getFullYear(), (date.getMonth()) + num, date.getDate(), date.getHours(), date.getMinutes(), date.getSeconds());
			case 'y':
				return new Date((date.getFullYear() + num), date.getMonth(), date.getDate(), date.getHours(), date.getMinutes(), date.getSeconds());
		}
	}

	/**
	 * 比较日期差 dtEnd 格式为日期型或者有效日期格式字符串
	 * @param strInterval string  可选值 y 年 m月 d日 w星期 ww周 h时 n分 s秒  
	 * @param dtStart Date  可选值 y 年 m月 d日 w星期 ww周 h时 n分 s秒
	 * @param dtEnd Date  可选值 y 年 m月 d日 w星期 ww周 h时 n分 s秒 
	 */
	this.dateDiff = function(strInterval, dtStart, dtEnd) {
		switch (strInterval) {
			case 's':
				return parseInt((dtEnd - dtStart) / 1000);
			case 'n':
				return parseInt((dtEnd - dtStart) / 60000);
			case 'h':
				return parseInt((dtEnd - dtStart) / 3600000);
			case 'd':
				return parseInt((dtEnd - dtStart) / 86400000);
			case 'w':
				return parseInt((dtEnd - dtStart) / (86400000 * 7));
			case 'm':
				return (dtEnd.getMonth() + 1) + ((dtEnd.getFullYear() - dtStart.getFullYear()) * 12) - (dtStart.getMonth() + 1);
			case 'y':
				return dtEnd.getFullYear() - dtStart.getFullYear();
		}
	}

	/**
	 * 字符串转换为日期对象
	 * @param date Date 格式为yyyy-MM-dd HH:mm:ss，必须按年月日时分秒的顺序，中间分隔符不限制
	 */
	this.strToDate = function(dateStr) {
		var data = dateStr;
		var reCat = /(\d{1,4})/gm;
		var t = data.match(reCat);
		t[1] = t[1] - 1;
		eval('var d = new Date(' + t.join(',') + ');');
		return d;
	}

	/**
	 * 把指定格式的字符串转换为日期对象yyyy-MM-dd HH:mm:ss
	 * 
	 */
	this.strFormatToDate = function(formatStr, dateStr) {
		var year = 0;
		var start = -1;
		var len = dateStr.length;
		if ((start = formatStr.indexOf('yyyy')) > -1 && start < len) {
			year = dateStr.substr(start, 4);
		}
		var month = 0;
		if ((start = formatStr.indexOf('MM')) > -1 && start < len) {
			month = parseInt(dateStr.substr(start, 2)) - 1;
		}
		var day = 0;
		if ((start = formatStr.indexOf('dd')) > -1 && start < len) {
			day = parseInt(dateStr.substr(start, 2));
		}
		var hour = 0;
		if (((start = formatStr.indexOf('HH')) > -1 || (start = formatStr.indexOf('hh')) > 1) && start < len) {
			hour = parseInt(dateStr.substr(start, 2));
		}
		var minute = 0;
		if ((start = formatStr.indexOf('mm')) > -1 && start < len) {
			minute = dateStr.substr(start, 2);
		}
		var second = 0;
		if ((start = formatStr.indexOf('ss')) > -1 && start < len) {
			second = dateStr.substr(start, 2);
		}
		return new Date(year, month, day, hour, minute, second);
	}

	/**
	 * 日期对象转换为毫秒数
	 */
	this.dateToLong = function(date) {
		return date.getTime();
	}

	/**
	 * 毫秒转换为日期对象
	 * @param dateVal number 日期的毫秒数 
	 */
	this.longToDate = function(dateVal) {
		return new Date(dateVal);
	}

	/**
	 * 判断字符串是否为日期格式
	 * @param str string 字符串
	 * @param formatStr string 日期格式， 如下 yyyy-MM-dd
	 */
	this.isDate = function(str, formatStr) {
		if (formatStr == null) {
			formatStr = "yyyyMMdd";
		}
		var yIndex = formatStr.indexOf("yyyy");
		if (yIndex == -1) {
			return false;
		}
		var year = str.substring(yIndex, yIndex + 4);
		var mIndex = formatStr.indexOf("MM");
		if (mIndex == -1) {
			return false;
		}
		var month = str.substring(mIndex, mIndex + 2);
		var dIndex = formatStr.indexOf("dd");
		if (dIndex == -1) {
			return false;
		}
		var day = str.substring(dIndex, dIndex + 2);
		if (!this.isNumber(year) || year > "2100" || year < "1900") {
			return false;
		}
		if (!this.isNumber(month) || month > "12" || month < "01") {
			return false;
		}
		if (day > this.getMaxDay(year, month) || day < "01") {
			return false;
		}
		return true;
	}

	/**
	 *	判断某年某月的天数
	 */
	this.getMaxDay = function(year, month) {
		if (month == 4 || month == 6 || month == 9 || month == 11)
			return "30";
		if (month == 2)
			if (year % 4 == 0 && year % 100 != 0 || year % 400 == 0)
				return "29";
			else
				return "28";
		return "31";
	}

	/**
	 * 把日期分割成数组 [年、月、日、时、分、秒]
	 */
	this.toArray = function(myDate) {
		myDate = arguments[0] || new Date();
		var myArray = Array();
		myArray[0] = myDate.getFullYear();
		myArray[1] = myDate.getMonth() + 1;
		myArray[2] = myDate.getDate();
		myArray[3] = myDate.getHours();
		myArray[4] = myDate.getMinutes();
		myArray[5] = myDate.getSeconds();
		return myArray;
	}

	/**
	 * 取得日期数据信息  
	 * 参数 interval 表示数据类型  
	 * y 年 M月 d日 w星期 w周 h时 n分 s秒  
	 */
	this.datePart = function(interval, myDate) {
		myDate = arguments[1] || new Date();
		var partStr = '';
		var Week = ['日', '一', '二', '三', '四', '五', '六'];
		switch (interval) {
			case 'y':
				partStr = myDate.getFullYear();
				break;
			case 'M':
				partStr = myDate.getMonth() + 1;
				break;
			case 'd':
				partStr = myDate.getDate();
				break;
			case 'w':
				partStr = Week[myDate.getDay()];
				break;
			case 'ww':
				partStr = myDate.WeekNumOfYear();
				break;
			case 'h':
				partStr = myDate.getHours();
				break;
			case 'm':
				partStr = myDate.getMinutes();
				break;
			case 's':
				partStr = myDate.getSeconds();
				break;
		}
		return partStr;
	}

	/**
	 * 取得当前日期所在月的最大天数  
	 */
	this.maxDayOfDate = function(date) {
		date = arguments[0] || new Date();
		date.setDate(1);
		date.setMonth(date.getMonth() + 1);
		var time = date.getTime() - 24 * 60 * 60 * 1000;
		var newDate = new Date(time);
		return newDate.getDate();
	}

	return this;
}

/**
 * 发送Ajax请求  
 * @param option (onSuccess, url, param, loading_Id)
 */
function request(option) {
    var ajaxRequest = creatAjaxRequest();
    if (ajaxRequest == null) {
        alert("您的浏览器不支持AJAX！");
        return;
    }
    ajaxRequest.onreadystatechange = function () {
        if (ajaxRequest.readyState == 4) {
            if (ajaxRequest.status >= 200 && ajaxRequest.status < 300 || ajaxRequest.status == 304) {
                option.onSuccess(ajaxRequest.responseText);
            }
        }
        else {
            if (!option.loading_Id)
                document.getElementById(option.loading_Id).innerHTML = "<div class='hook_con'><img class='loading_pic' src='images/loading.gif'/></div>";
        }
    };
    ajaxRequest.open("post", option.url, true);
    ajaxRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    ajaxRequest.send(option.param);
}

/**
 * 创建一个ajaxRequest对象
 */
function creatAjaxRequest() {
    var xmlHttp = null;
    if (window.XMLHttpRequest) {
        xmlHttp = new XMLHttpRequest();
    } else {
        try {
            xmlHttp = new ActiveXObject("Msxml2.XMLHTTP");
        } catch (e) {
            try {
                xmlHttp = new ActiveXObject("Microsoft.XMLHTTP");
            } catch (e) {
            }
        }
    }
    return xmlHttp;
}
