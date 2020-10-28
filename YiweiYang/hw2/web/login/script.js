var ridCompleted = false;
var uidCompleted = false;
function checkRidInput()
{
	var rid = document.getElementById("room_id").value;
  if (rid.length == 0) {
    if (ridCompleted == true){
        document.getElementById("r_info").style.display = "none";
     }
    ridCompleted = false;
  } else {
    if (ridCompleted == false){
        document.getElementById("r_info").style.display = "inline";
     }
    ridCompleted = true;
  }
  if (ridCompleted && uidCompleted){
      document.getElementById("login_button").style.background = "#5286ed";
      }else{
        document.getElementById("login_button").style.background = "#c8c8c8";
      }
}
function checkUidInput()
{
	var uid = document.getElementById("user_id").value;
  if (uid.length == 0) {
        if (uidCompleted == true){
        document.getElementById("u_info").style.display = "none";
     }
    uidCompleted = false;
  } else {
    if (uidCompleted == false){
        document.getElementById("u_info").style.display = "inline";
     }
    uidCompleted = true;
  }
  if (ridCompleted && uidCompleted){
      document.getElementById("login_button").style.background = "#5286ed";
    document.getElementById("login_button").disabled = false;
      }else{
        document.getElementById("login_button").style.background = "#c8c8c8";
      }
}

function submitInfo()
{
  var rid = document.getElementById("room_id").value;
  var uid = document.getElementById("user_id").value;
  if (isValid(rid)&&isValid(uid)){
      alert("成功！");
      }else{
    alert("失败，必须同时包含且仅包含[字母][数字][下划线]！");
  }
}

function isValid(str) 
{ 
  return /(?=.*[A-Za-z])(?=.*[\d])(?=.*_)(^[0-9a-zA-Z_]{1,}$)/.test(str); 
}