var rooturl = 'http://localhost:7890/';
var todo = angular.module('todo',['ngCookies']);

todo.config(function($interpolateProvider){
	$interpolateProvider.startSymbol('[[');
	$interpolateProvider.endSymbol(']]');

});

//初始化页面信息，时间采用当前日期
function initPage($scope){
	var date = new Date().Format("yyyy-MM-dd");
	$scope.Starttime = date;
	$scope.today = date;
	$scope.Level='0';
}

//获取后台的当天的信息
function getTodayTodos($scope,$http){
	var url = rooturl + 'todo/list/'+$scope.userid;
	
	$http.get(url).success(function(todos){
		$scope.loaded = true;
		$scope.todos = todos;
		//alert(todos);
		for(var i=0;i<todos.length;i++){
			//alert(todos[i].Title);
		}
	}).error(function(err){
		alert(err);
	});
}

//获取完成信息，前七天的数据
function getFinishTodos($scope,$http){
	var url = rooturl + 'todo/finishlist/'+$scope.userid;
	
	$http.get(url).success(function(todos){
		$scope.loaded = true;
		$scope.ftodos = todos;
		
	}).error(function(err){
		alert(err);
	});
}

todo.controller('TodoCtrl',function($scope,$http,$cookies){
	
	$scope.username = $cookies.username;
	$scope.userid = $cookies.userid;
	
	initPage($scope);
	
	//获取当天的未完成的工作 
	getTodayTodos($scope,$http);
	
	getFinishTodos($scope,$http);
	
	$scope.todayTodos=[];
	
	//点击增加按钮时,提交后台保存
	$scope.addTodo=function(){
		var todo = {Starttime:$scope.Starttime,Title:$scope.Title,Level:$scope.Level,Isfinish:0,UserId:$scope.userid};
		var urladd = rooturl + "todo/add";
		var postCfg = {
                headers: { 'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'},
                transformRequest: function(data){
					return $.param(data);
				}
            };
		$http.post(urladd,todo,postCfg).success(function(data){
			if(data.Code==0){
				alert(data.Msg);
			}else{
				todo = data.Data;
				$scope.todos.push(todo);
			}
			
		}).error(function(err){
			alert("连接服务器失败！");
		});
		
	};
	
	
	//点击完成按钮,提交后台，修改完成状态，再把内容划横线，去掉完成按钮
	$scope.todoFinish=function(todo){
		var urlf = rooturl + "todo/finish/"+todo.Id;
		$http.get(urlf).success(function(data){
			if(data.Code==0){
				alert(data.Msg);
			}else{
				todo.Isfinish=1;
			}
			
		}).error(function(err){
			alert("连接服务器失败！");
		});
	};
	
	//根据值，转化成对应的label
	$scope.Select2LabelText=function(val){
		
		var result = "";
		switch(val){
		case 0:
			result = "一般";
			break;
		case 1:
			result ="重要";break;
		case 2:
			result = "紧急";break;
		default:
			result = "一般";
		}
		return result;
	};
	$scope.Select2LabelClass=function(val){
		var result = "";
		switch(val){
		case 0:
			result = "label-primary";break;
		case 1:
			result ="label-warning";break;
		case 2:
			result = "label-danger";break;
		default:
			result = "label-primary";
		}
		return result;
	};
	
	$scope.FormatTime=function(time){
		return time.substr(0,10);
	};
		
	
});




//日期格式化
Date.prototype.Format=function(fmt){
	var o={
		"M+":this.getMonth()+1,
		"d+":this.getDate(),
		"h+":this.getHours(),
		"m+":this.getMinutes(),
		"s+":this.getSeconds(),
		"q+":Math.floor((this.getMonth()+3)/3),//季度
		"S":this.getMilliseconds()//毫秒
	};
	if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
    for (var k in o)
    	if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
    return fmt;

	
	return fmt;
	
}


$.fn.datepicker.defaults.format = "yyyy-MM-dd HH:mm:ss";
$.fn.datepicker.dates['zh-CN'] = {
		days: ["星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"],
		daysShort: ["周日", "周一", "周二", "周三", "周四", "周五", "周六", "周日"],
		daysMin:  ["日", "一", "二", "三", "四", "五", "六", "日"],
		months: ["01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"],
		monthsShort: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
		today: "今日",
		//format: "yyyy年mm月dd日",
		weekStart: 1
	};
$.fn.datepicker.defaults.language = "zh-CN";
