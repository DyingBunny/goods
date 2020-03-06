;(function($){
	
	/**
	 * radiobutton  赋值
	 */	
	$.fn.initRadioBox = function(label_val){
		
	    this.find("label").each(function(){
	    	
			if($(this).html() == label_val){
				$(this).prev().addClass("checked");
				$(this).prev().children("input").addClass("checked");
				$(this).prev().children("input").attr("checked", "checked");
				$(this).prev().siblings(".span-radio").removeClass("checked");
				$(this).prev().siblings(".span-radio").find("input").removeAttr("checked"); 
				
				return;
			}				
		});
	}	
    
	/**
	 * radiobutton 点击
	 */
	$.fn.handleRadioClick = function(){
		
		  if (this.hasClass("checked")) {
			  
//			  this.removeClass("checked");
//			  this.children().removeAttr("checked");	
			  
	      } else {
	    	  
	    	  this.addClass("checked");
	    	  this.children().attr("checked", "checked");	    
	    	  this.siblings(".span-radio").removeClass("checked");
	    	  this.siblings(".span-radio").find("input").removeAttr("checked");          
	      }
	}
	/**
	 * checkbox 初始化
	 */
    $.fn.initCheckBox = function(way,check_val){
		switch(way){
		case "id":
			
			var target_dom=$("#"+check_val);
			target_dom.addClass("checked")
	        target_dom.children("input").addClass("checked");
			target_dom.children("input").attr("checked", "checked"); 
			
			break;
			
        case "code":
			
        	 this.find(".span-checkbox").each(function(){
			    	
					if($(this).attr('code') == check_val){
						$(this).addClass("checked");
						$(this).children("input").attr("checked", "checked");
						return;
					}				
			 });
        	 break;
			
		case "lable":
			  this.find("label").each(function(){
			    	
					if($(this).html() == check_val){
						$(this).prev().addClass("checked");
						$(this).prev().children("input").addClass("checked");
						$(this).prev().children("input").attr("checked", "checked"); 
						
						return;
					}				
				});
			break;
		}
	
	}	
	/**
	 * checkbox 点击
	 */
	$.fn.handleCheckBoxClick = function(){
		
		 if(this.hasClass("checked")) {			 
             this.removeClass("checked");
             this.children().removeAttr("checked");
             
         }else{
             this.addClass("checked");
             this.children().attr("checked", "checked");          
         }	
	}
	
	$.fn.handleCheckBoxNoState = function(){           /* add by sunfang 2016-12-05*/
        this.removeClass("checked");
        this.children().removeAttr("checked");
	}
	
	/**
	 * checkbox全选 点击
	 * @list_dom_str 全选作用的checkbox区域的父节点的dom 例如 #service_table, .service-table
	 */
	$.fn.handleALLCheckClick = function(list_dom_str){
		var _this=this.get(0);
		if(_this.checked){	
			$(list_dom_str+" :checkbox").prop("checked","checked");
			$(list_dom_str+" .span-checkbox").addClass("checked");
			this.parent(".span-checkbox").addClass("checked");
		}else{		
			$(list_dom_str+" :checkbox").removeAttr("checked");
			$(list_dom_str+" .span-checkbox").removeClass("checked");
			this.parent(".span-checkbox").removeClass("checked");
		}
	}
	/**
	 * 初始化左侧导航
	 */
	$.fn.initLeftNav = function(name){		
		var navdom=this.find('li a');		
		$.each(navdom,function(index,item){
			if(item.innerHTML==name || (typeof($(item).attr('code')) != "undefined" && $(item).attr('code') == name)){
				$(this).parent('li').addClass('active');
				$(this).parents('div.collapse').addClass('in');
				$(this).parents('div.collapse').prev('.left-nav-tile').attr('aria-expanded','true');
				$(this).parents('div.collapse').prev('.left-nav-tile').removeClass('collapsed');
				return false;
			}
		});
	}
	/**
	 * 高度适应
	 * @target 目标dom
	 */
	$.fn.handleSimilarHeight = function(target){
		var target_height=$(target).height();
		if(this.height() < target_height)
		  this.css('min-height',target_height+'px');
	}
	/**
	 * 高度设置
	 * @target 目标元素高度
	 * @redupx 被-掉的px值
	 */
	$.fn.handleMinHeight = function(redupx){
		var target_height=$(window).height();
		this.css('min-height',(target_height-redupx)+'px');
	}
	/**
	 * 筛选条目,初始化
	 */
	$.fn.initActiveFilter = function(filter){
		if(filter == null)
			this.find("a.item-all").addClass("active");
		else{
			this.find("#"+filter).addClass("active");	         
			this.find("#"+filter).append('<i class="remove-icon"></i>');
		}
	}
	
	/**
	 * 重新绘制筛选框
	 */
	 $.fn.reDrewFilter = function(unfold){
		    var f_top=this.find('a:first-child').offset();
		    var l_top=this.find('a:last-child').offset();
		    
		    if(l_top.top > f_top.top){
		    	var more_html='<dd class="pull-left show-more"><h3><i class="fold-icon"></i><span>更多</span></h3></dd>';
				if(parseInt(unfold)==1){
	               more_html='<dd class="pull-left show-more unfold"><h3><i class="unfold-icon"></i><span>收起</span></h3></dd>';
	               this.css('height','auto');
				}			
				this.parent('dl').append(more_html);
		    }   	
	   }   
	
	/**
	 * 技术介绍标签置顶固定
	 */
	$.fn.smartFloat = function() {
		var position = function(element) {
			var top = element.position().top, pos = element.css("position");
			$(window).scroll(function() {
				var scrolls = $(this).scrollTop();
				if (scrolls > top) {
					if (window.XMLHttpRequest) {
						element.css({
							position: "fixed",
							top: 0
						}); 
					} else {
						element.css({
							top: scrolls
						}); 
					}
				}else {
					element.css({
						position: pos,
						top: top
					}); 
				}
			});
		};
		return $(this).each(function() {
			position($(this));      
		});
	}
	
	//$('input, textarea').placeholder();
	
})(jQuery); 

window.console = window.console || (function(){ 
	var c = {}; c.log = c.warn = c.debug = c.info = c.error = c.time = c.dir = c.profile 
	= c.clear = c.exception = c.trace = c.assert = function(){}; 
	return c; 
})();