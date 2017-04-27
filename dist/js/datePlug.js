//体检预约日期插件
	var DataCheck=false;//是否已经预约状态
	var DataYear=0;//预约的年份
	var DataMonth=0;//预约的月份
	var DataDay=0;//预约的日子
	var dateReady=false;//日期状态是否完成
	$(".myDate td").css("height",$(".myDate td").width()+"px");
	//ie兼容处理
	function NewDate(str) { 
		str = str.split('-');
		var date = new Date(); 
		date.setUTCFullYear(str[0], str[1] - 1, str[2]);
		date.setUTCHours(0, 0, 0, 0); 
		return date;
	}
	//闰年判断
	function isLeapYear (Year) {
		if (((Year % 4)==0) && ((Year % 100)!=0) || ((Year % 400)==0)) {
		return (true);
		} else { return (false); }
	};
	var capatityed={};//客满
	var parts={};//项目满
	var myTime=new Date();
	var year=myTime.getFullYear();
	var month=myTime.getMonth()+1;//要显示的月份,默认显示当前月
	var rightYear=year;//当前年
	var rightDay=myTime.getDate();//当前日
	var rightMon=month;//当前月
	var offdays={};//休息日
	
		
	//月份所有天显示
	
	function showDay(year,month){
		var day1=NewDate(year+"-"+month+"-"+"01").getDay();
		
		$(".myDate>h3.dateTitle>a").html(year+"年"+month+"月");
		$(".myDate tbody td").each(function(){
			$(this).removeAttr("class").html("");
		});
		//月份天数判断
		if(month==2){
			//2月份
			
			if(isLeapYear(year)){
				var maxDay=29;
			}else{
				var maxDay=28;
			}
			
		}else if([1,3,5,7,8,10,12].indexOf(month)>=0){
			var maxDay=31;
		}else{
			var maxDay=30;
		}

		var monthday=1;
		//行数
		for(var line=0;line<6;line++){
			if(line==0){
				//首行
				for(var i=day1;i<7;i++){
					$(".myDate tbody>tr:first-of-type>td").eq(i).html(monthday);
					dayState();
					monthday++;
				}
			}else{
				for(var i=0;i<7;i++){
					$(".myDate tbody>tr").eq(line).children("td").eq(i).html(monthday);
					dayState();
					monthday++;
					if(monthday>maxDay){
						return 0;
					}
					
				}
			}
		}
		//当天状态判读
		function dayState(){
			//checkmonth匹配后台数据的月份格式
			var checkmonth=year+((month<10)?("0"+month):string(month));

			//是否当天
			$(".myDate tbody>tr").eq(line).children("td").eq(i).removeAttr("class").addClass("permit");
			//是否星期休息日
			if(offdays[checkmonth]){
				if(offdays[checkmonth].indexOf(monthday)==-1){
					$(".myDate tbody>tr").eq(line).children("td").eq(i).removeClass("reset");
				}else{
					$(".myDate tbody>tr").eq(line).children("td").eq(i).removeAttr("class").addClass("reset");
				}
			};
			
			//是否客满
			if(capatityed[checkmonth]){
				if(capatityed[checkmonth].indexOf(monthday)==-1){
					$(".myDate tbody>tr").eq(line).children("td").eq(i).removeClass("full");		
				}else{
					$(".myDate tbody>tr").eq(line).children("td").eq(i).removeAttr("class").addClass("full");			
				}
			}
			
			//项目已满
			if(parts[checkmonth]){
				if(parts[checkmonth].indexOf(monthday)==-1){
					$(".myDate tbody>tr").eq(line).children("td").eq(i).removeClass("part");		
				}else{
					$(".myDate tbody>tr").eq(line).children("td").eq(i).removeAttr("class").addClass("part");			
				}
			}
			//日期范围之外的monthday状态
			if(year==rightYear){
				if(((month==rightMon)&&(monthday<=rightDay))||((month==(rightMon+2))&&(monthday>rightDay))){
					$(".myDate tbody>tr").eq(line).children("td").eq(i).removeAttr("class").addClass("forbid");
				}
			}else if(year==(rightYear+1)){
				if(((month+2)==(rightMon+12))&&(monthday>rightDay)){
					$(".myDate tbody>tr").eq(line).children("td").eq(i).removeAttr("class").addClass("forbid");
				}
			}
			//选中
			if((DataCheck)&&(DataYear==year)&&(DataMonth==month)&&(DataDay==monthday)){
				$(".myDate tbody>tr").eq(line).children("td").eq(i).removeAttr("class").addClass("checked");			
			}
		};
		
	};
	
	showDay(year,month);
	
	//左右点击切换月份事件
	$(".myDate>h3.dateTitle>span").click(function(e){
		e.stopPropagation();
		if($(this).index()==2){
			if(month+1>(rightMon+2)){
				return 0;
			}
			month++;
			if(month>12){
				year++;
				month=1;
			}
			showDay(year,month);
		}else if($(this).index()==0){
			if((month-1)<rightMon){
				return 0;
			}
			month--;
			if(month<1){
				year--;
				month=12;
			}
			showDay(year,month);
				
		}
	});
	//预约日期选择
	$(".myDate tbody td").click(function(e){
		if(dateReady){
			e.stopPropagation();
			if($(this).hasClass("permit")){
				$(this).parents("tbody").find(".checked").removeClass("checked").addClass("permit");
				$(this).removeAttr("class").addClass("checked");
				$(".myDate caption>div").css("color","red").html("已选中"+year+"年"+month+"月"+$(this).text()+"日");
				DataYear=year;
				DataMonth=month;
				DataDay=$(this).html();
				return DataCheck=true;
			}
			if(($(this).hasClass("checked"))&&(DataCheck)){
				$(this).removeAttr("class").addClass("permit");
				$(".myDate caption>div").removeAttr("style").html("<p>可以预约</p><p>不可预约</p>");
				DataYear=0;
				DataMonth=0;
				DataDay=0;
				return DataCheck=false;
			}
		}else{
			mui.toast("请先选择分院");
		}
	});




