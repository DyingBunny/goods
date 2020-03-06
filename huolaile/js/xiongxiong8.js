var body_width = 0; //可用区域的宽度 (浏览器内)
var body_height = 0; //可用区域的高度 (浏览器内)

var navbar_top_height = 0; //顶部菜单的高度
var navbar_bottom_height = 0; //底部菜单的高度

var index_guanggao_height = 0; //首页广告的高度
var index_guanggao_height_min = 100; //首页广告的高度
var index_guanggao_height_max = 200; //首页广告的高度

var container_height = 0;
var index_guanggao_height = 0;
var index_buttons_height = 0;
var index_buttons_margin_top = 0;

var daisonghuo_height = 0; //首页代收货按钮的高度
var peisongzhong_height = 0; //首页配送中按钮的高度
var yiwancheng_height = 0; //首页已完成按钮的高度
var yiwancheng_margin_top = 0 ; //首页已完成按钮的margin-top边距
var loading_div_margin_top = 0; //loading div的边距

var navbar_top_ID = 'navbar-top';
var navbar_bottom_ID = 'navbar-bottom';
var container_ID = 'container';
var index_guanggao_ID = 'index_guanggao';
var index_buttons_ID = 'index_buttons';
var index_buttons3_ID = 'index_buttons3';
var daisonghuo_ID = 'daisonghuo';
var peisongzhong_ID = 'peisongzhong';
var yiwancheng_ID = 'yiwancheng';
var daichuli_ID = 'daichuli';
var loading_div_ID = 'loading_div';

$(function() {
	resizeAll();
	//窗口调整以后
	$(window).resize(function() {
		resizeAll();
	});
	
	shuaxinYemian(30); //定时刷新页面
});

/**
 * 重置所有参数
 */
function resizeAll() {
	getWidthHeight();
	setWidthHeight();
	
	LiuchengWenzi(); //展开流程的文字
}

/**
 * 计算出宽高及显示模式
 */
function getWidthHeight() {
    body_width = $(window).width(); //可用区域的宽度 (浏览器内)
    body_height = $(window).height(); //可用区域的高度 (浏览器内)
	
	navbar_top_height = $('#' + navbar_top_ID).height(); //顶部菜单的高度
	navbar_bottom_height = $('#' + navbar_bottom_ID).height(); //底部菜单的高度
	
	container_height = body_height - navbar_top_height - navbar_bottom_height; //内容容器的高度
	
	//广告的高度
	index_guanggao_height = container_height * 0.30;
	if (index_guanggao_height > index_guanggao_height_max) {
		index_guanggao_height = index_guanggao_height_max;
	} else if (index_guanggao_height < index_guanggao_height_min) {
		index_guanggao_height = index_guanggao_height_min;
	}
	
	index_buttons_height = container_height - index_guanggao_height;
	index_buttons_margin_top = index_buttons_height * 0.07 / 2;
	
	daisonghuo_height = index_buttons_height * 0.93; //首页代收货按钮的高度
	peisongzhong_height = index_buttons_height * 0.45; //首页配送中按钮的高度
	yiwancheng_height = index_buttons_height * 0.45; //首页已完成按钮的高度
	
	yiwancheng_margin_top = daisonghuo_height - peisongzhong_height - yiwancheng_height;
	
	index_buttons_height = index_buttons_height - index_buttons_margin_top; //在最后要减去margin-top的值
	
	loading_div_margin_top = container_height / 2 - 30;
	
	//console.log(yiwancheng_height);
}

/**
 * 给各个div设定尺寸
 */
function setWidthHeight() {
	if($('#' + index_guanggao_ID).length){
		$('#' + index_guanggao_ID).css("height", index_guanggao_height);
	}
	
	if($('#' + index_buttons_ID).length){
		$('#' + container_ID).css("height", container_height); //设置内容部分的容器高度
		$('#' + container_ID).css("background-color", "#FFFFFF"); //设置内容部分的背景色
		$('#' + index_buttons_ID).css("height", index_buttons_height); //首页按钮部分的高度
		$('#' + index_buttons_ID).css("margin-top", index_buttons_margin_top); //首页按钮部分的margin-top边距
		
		
		var index_button_text_m = $('#' + daisonghuo_ID).width() * 0.25 / 2 - 3; //算出按钮背景相关的偏移值
	
		if($('#' + index_buttons3_ID).length){
			$('#' + daisonghuo_ID).css("height", peisongzhong_height);
			$('#' + daichuli_ID).css("height", yiwancheng_height);
			$('#' + daichuli_ID).css("margin-top", yiwancheng_margin_top); //已完成的margin-top边距
			$('#' + peisongzhong_ID).css("height", peisongzhong_height);
			$('#' + yiwancheng_ID).css("height", yiwancheng_height);
			$('#' + yiwancheng_ID).css("margin-top", yiwancheng_margin_top); //已完成的margin-top边距
		
			$('#' + daisonghuo_ID + " .index_button_text").css("margin-top", (peisongzhong_height / 2) + index_button_text_m);
			$('#' + peisongzhong_ID + " .index_button_text").css("margin-top", (peisongzhong_height / 2) + index_button_text_m);
			$('#' + yiwancheng_ID + " .index_button_text").css("margin-top", (yiwancheng_height / 2) + index_button_text_m);
			$('#' + daichuli_ID + " .index_button_text").css("margin-top", (yiwancheng_height / 2) + index_button_text_m);
		} else {
			$('#' + daisonghuo_ID).css("height", daisonghuo_height);
			$('#' + peisongzhong_ID).css("height", peisongzhong_height);
			$('#' + yiwancheng_ID).css("height", yiwancheng_height);
			$('#' + yiwancheng_ID).css("margin-top", yiwancheng_margin_top); //已完成的margin-top边距
		
			$('#' + daisonghuo_ID + " .index_button_text").css("margin-top", (daisonghuo_height / 2) + index_button_text_m);
			$('#' + peisongzhong_ID + " .index_button_text").css("margin-top", (peisongzhong_height / 2) + index_button_text_m);
			$('#' + yiwancheng_ID + " .index_button_text").css("margin-top", (yiwancheng_height / 2) + index_button_text_m);
		}
	}
	
	$('#' + navbar_bottom_ID + ' a').css("background-size", navbar_bottom_height * 0.6); //底部菜单的ICO图片大小
	$('#' + loading_div_ID).css("margin-top", loading_div_margin_top);
}

//定时刷新页面
function shuaxinYemian(time) {
	var interval = setInterval(function() {
		$("#container").load(location.href+" #container>*","", function(){
			resizeAll();
		});
	}, time * 1000);

	$("input:text").click(function() {
		clearInterval(interval);
	});

	$("input:text").focus(function() {
		clearInterval(interval);
	});
}

//展开流程的文字
function LiuchengWenzi() {
	$("#liucheng_liucheng div.row").click(function(){
		$(this).find('.wenzi').removeClass("text-overflow");
	});
}

//显示loading，跳转到某个url
function gotoUrl(url, showImg){
	if (showImg) {
		$('#loading_main').css('display','block'); 
	}
	if (url) {
		window.open(url,'_self');
	}
	event.stopPropagation();
}