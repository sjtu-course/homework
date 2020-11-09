function isValid(str){ 
    // 正则表达式：两个//分别表示开始和结束, ^表示开始字符串, $表示结束字符串, \w表示包含【a-z，A-Z, _ , 0-9】, +表示一个或者多个\w
    return /^\w+$/.test(str);
}

function enter_room(){
    var room = document.getElementById("room_text").value;
    var user = document.getElementById("user_text").value;
    if (isValid(room) && isValid(user)){
        alert("成功");
    }
    else{
        alert("失败");
    }
}