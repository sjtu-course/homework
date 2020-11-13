function enter_room(){
    var room = document.getElementById("room_text").nodeValue;
    var user = document.getElementById("user_text").nodeValue;
    // document.getElementById("enter_button").removeAttribute("disabled", true);
    // document.getElementById("enter_button").style.color="#5286ed";
    
    if (room == "" || user == ""){
        document.getElementById("enter_button").setAttribute("disabled", true);
        alert("成功");
    }
    else{
        // document.getElementById("enter_button").removeAttribute("disabled", true);
        document.getElementById("enter_button").style.backgroundColor = "#5286ed";
        // document.getElementById('enter_button').style.border = '1px #5286ed solid'
        
    }
}