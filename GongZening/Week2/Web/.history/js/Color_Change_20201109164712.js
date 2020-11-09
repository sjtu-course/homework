function color_change(){
    var room = document.getElementById("room_text").value;
    var user = document.getElementById("user_text").value;
    if (room == "" || user == ""){
        // document.getElementById("enter_button").setAttribute("disabled", true);
    }
    else{
        // document.getElementById("enter_button").removeAttribute("disabled", true);
        document.getElementById("enter_button").style.backgroundColor = "#5286ed";
        // document.getElementById('enter_button').style.border = '1px #5286ed solid'
    }
}