/* $(document).ready(function(){
    var username,password;
    var r_data = {
        "username":username,
        "password":password
    };
    var login_option = {
        url: 'http://175.24.22.208:80/auth',
        requestData: r_data,
		beforeSend: function(request) {
            request.setRequestHeader("Authorization", token);
        },
        successHandler: function(data){
            if(data.code == "000000"){
                window.location="/views/admin/centre/basicInformation.jsp";
            }else if(data.code == "999999"){
                Confirm.show('��ʾ',"warning", "�˺Ż��������", {
                    'ȷ��': {
                        'callback': function() {
                            Confirm.hide();
                        }
                    }
                });
            }
        },
        errorHandler: function(data){
        }
    };
    document.onkeydown = function(e){
        var ev = document.all ? window.event : e;
        if(ev.keyCode==13) {
            var username=$("#username").val();
            var password=$("#password").val();
            if(username !='' && password !=''){
                r_data.username = username;
                r_data.password = password;
                login_option.requestData = r_data;
                ajaxCallScript(login_option);
            }else{
                Confirm.show('��ʾ',"warning", "��¼��Ϣ��д������", {
                    'ȷ��': {'callback': function() {Confirm.hide();}}
                })
            }
        }
    };
    document.onkeyup = function(e){
        var ev = document.all ? window.event : e;
        if(ev.keyCode==13) {
            var username=$("#username").val();
            var password=$("#password").val();
            if(username !='' && password !=''){
                r_data.username = username;
                r_data.password = password;
                login_option.requestData = r_data;
                ajaxCallScript(login_option);
            }else{
                Confirm.show('��ʾ',"warning", "��¼��Ϣ��д������", {
                    'ȷ��': {'callback': function() {Confirm.hide();}}
                })
            }
        }
    };
    $(".loginBtn").on("click",function(){		alert("11")
        var username=$("#username").val();
        var password=$("#password").val();
        if(username !='' && password !=''){
            r_data.username = username;
            r_data.password = password;
            login_option.requestData = r_data;
            ajaxCallScript(login_option);

		
			
			
			
$.ajax({
		url:"175.24.22.208:80/auth",
        contentType: "application/json",
        dataType: "json",
        data:{ 
			"username":"qianzihang",
			"password":"123"
	
		},
        cache: false,
        processData: false,
        async:status,
        type:"put",
        success:function (res) {
			console.log(res)
			
			
        },
        error:function(){
        },
        
    });

        }else{
			alert("��¼��Ϣ������")
			
            /*Confirm.show('��ʾ',"warning", "��¼��Ϣ������", {
                'ȷ��': {'callback': function() {Confirm.hide();}}
            })
        }
    });
}) */

$(document).ready(function(){
	
	$(".loginBtn").on("click",function(){	
		var username=$("#username").val();
		var password=$("#password").val();
		console.log(username+"=="+password)
		
		
		  $.ajax({
			 
			  url: "http://175.24.22.208:80/auth",//�����ַaaa
			  contentType: "application/json",
			   type: "POST",
			  data: {
		        username:"qianzihang",
		        password:"123"
		     },
			  success: (data) => {
				if (data.return_code === "SUCCESS") {
				  this.deviceList = data.return_info;
				  // console.log(this.deviceList);
				} else {
				  alert("������󣡴�����Ϣ��" + data.return_msg);
				}
			  },
			  error: (err) => {
				this.loginLoading = false;
				alert("������Ϣ��" + JSON.stringify(err))
			  }
			})
		
	})
	
	
	
	
	/* tab�л� */
	$("#switch a").click(function(){
		var idName = $(this).attr("class");
		var index = $(this).index();
		if(index == 0){
			$("#switch a").eq(1).attr("class","switch_btn");
			$(this).attr("class","switch_btn_focus");
			$("#switch_bottom").css("left","0");
			$("#web_qr_login").css("display","block");
			$("#qlogin").css("display","none");
		}else if(index == 1){
			$("#switch a").eq(0).attr("class","switch_btn");
			$(this).attr("class","switch_btn_focus");
			$("#switch_bottom").css("left","154px");
			$("#web_qr_login").css("display","none");
			$("#qlogin").css("display","block");
		}
		
		
	})
	
	
})

