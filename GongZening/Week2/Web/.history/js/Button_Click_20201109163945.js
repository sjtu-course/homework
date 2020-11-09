function enter_room(){
    var room = document.getElementById("room_text").value;
    var user = document.getElementById("user_text").value;
    // document.getElementById("enter_button").style.color="#5286ed";
    
    if (room == null || user == null){
        // document.getElementById("enter_button").setAttribute("disabled", true);
        alert(room);
    }
    else{
        // document.getElementById("enter_button").removeAttribute("disabled", true);

        // document.getElementById("enter_button").removeAttribute("disabled", true);
        document.getElementById("enter_button").style.backgroundColor = "#5286ed";
        // document.getElementById('enter_button').style.border = '1px #5286ed solid'
        
    }
}