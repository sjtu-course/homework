function enter_room(){
    var room = document.getElementById("room_text").nodeValue;
    var user = document.getElementById("user_text").nodeValue;
    if (room == "" || user == ""){
        document.getElementById("enter_button").setAttribute("disabled", true);
    }
    else{
        // document.getElementById("enter_button").removeAttribute("disabled", true);
        document.getElementById("enter_button").style.backgroundColor="#5286ed";
    }
}