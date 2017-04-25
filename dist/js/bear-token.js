var bearToken="";
var patt = new RegExp("bear_token=");
function access(){	
	//获取token
	if(window.location.search.length>0){
		bearToken=window.location.search.replace("?bear_token=","");
		localStorage.bearToken=bearToken;
	}
	//读取token
	if(localStorage.bearToken){
		bearToken=localStorage.bearToken;		
	}else{
		mui.alert("不合法请求,需要微信登陆");
	}
	//ajax请求设置
	$.ajaxSetup({
		headers:{"Beartoken":bearToken},
		contentType:"application/json",
		beforeSend:function(){
			mui.toast("加载中");
		}
	});	
		
};
