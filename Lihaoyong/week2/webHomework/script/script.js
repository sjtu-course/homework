(
    function () {
        let roomID = document.getElementById("roomID")
        let userID = document.getElementById("userID")
        let enter = document.getElementById("enter")
  
        function checkInput() {
            if(roomID.value != "" && userID.value != ""){
                enter.style.backgroundColor="#5286ed"
                enter.removeAttribute("disabled")
            }else{
                enter.style.backgroundColor="#c8c8c8"
                enter.setAttribute("disabled", true)
            }
        }
  
        function eventClick() {
          // alert("成功")
          var success = /^([a-zA-Z]|[0-9]|_)+$/.test(roomID.value + userID.value) ? "成功" : "失败"
          alert(success)
        }
  
        enter.addEventListener("click", eventClick)
        roomID.addEventListener("input", checkInput)
        userID.addEventListener("input", checkInput)
    }
  )()