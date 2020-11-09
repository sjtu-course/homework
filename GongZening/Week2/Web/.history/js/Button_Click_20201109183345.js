function isValid(str){ 
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