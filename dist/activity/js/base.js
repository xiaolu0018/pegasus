var baseUrl="http://hd1.dahe100.cn";
var weLink="http://mp.weixin.qq.com/s?__biz=MzAwNDE4OTgyNw==&mid=506735864&idx=1&sn=ccb723be42baa87a748a4f511a2f8260&chksm=00ef37143798be02f90644c35eb2a0e95e4b38b6c181c0711506d0ca8874039328a0b46c9fbf&scene=18#wechat_redirect";
//提示框显示
var $toast = $('#toast');
function showToast(htdata,myclass){
	if(myclass){
		$("#toast i").removeAttr("class").addClass(myclass);
	}else{
		$("#toast i").removeAttr("class").addClass("weui_icon_toast");
	}
    if ($toast.css('display') != 'none') return;
	$("#toast>.weui_toast>p").html(htdata);
    $toast.fadeIn(100);
    setTimeout(function () {
         $toast.fadeOut(100);
    }, 2000);
};
//toast提示符
var myinfo="weui_icon_info_circle";
var mywarn="weui_icon_warn";
//alert框显示
function showAlert(aldata){
	$("#alertCon").html(aldata);
	$("#myAlert").show();
	$("#myAlert").on('touchstart','a',function(){
		$("#alertCon").html("");
		$("#myAlert").hide();
	});
};
//获取url参数
var request={};
function UrlSearch(){
   var name,value; 
   var str=location.href; //取得整个地址栏
   var num=str.indexOf("?") 
   str=str.substr(num+1); //取得所有参数   stringvar.substr(start [, length ]

   var arr=str.split("&"); //各个参数放到数组里
   for(var i=0;i < arr.length;i++){ 
	    num=arr[i].indexOf("="); 
	    if(num>0){ 
		    name=arr[i].substring(0,num);
		    value=arr[i].substr(num+1);
		    request[name]=value;		    
	    } 
    }
   if(request.name){
   		request.name=decodeURIComponent(request.name);
   }
   if(request.company){
   		request.company=decodeURIComponent(request.company);
   }
   
}