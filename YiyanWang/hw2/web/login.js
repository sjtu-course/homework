(
  function () {
      let room = document.getElementById("roomID")
      let name = document.getElementById("name")
      let btn = document.getElementById("btn")

      function checkInput() {
          if(room.value != "" && name.value != ""){
              btn.style.backgroundColor="#5286ed"
              btn.removeAttribute("disabled")
          }else{
              btn.style.backgroundColor="#c8c8c8"
              btn.setAttribute("disabled", true)
          }
      }

      function eventClick() {
        // alert("成功")
        var success = /^([a-zA-Z]|[0-9]|_)+$/.test(room.value + name.value) ? "成功" : "失败"
        alert(success)
      }

      btn.addEventListener("click", eventClick)
      room.addEventListener("input", checkInput)
      name.addEventListener("input", checkInput)
  }
)()