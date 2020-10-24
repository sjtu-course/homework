(
    function () {
        // 获取对应的元素
        let room = document.getElementById("room")
        let username = document.getElementById("username")
        let btn = document.getElementById("submit")
      
        function handleClick() {
            let room_value = room.value 
            let username_value = username.value 
            let regexp = /^([a-zA-Z]|[0-9]|_)+$/
            if(regexp.test(room_value) && regexp.test(username_value)){
                alert("成功")
            }else{
                alert("失败")
            }
        }

        function validate() {
            // 获取输入框中的值
            let room_value = room.value 
            let username_value = username.value 
            // 判断是否均有输入
            if(room_value && username_value){
                btn.style.backgroundColor="#5286ed"
                btn.removeAttribute("disabled")
            }else{
                btn.style.backgroundColor="#c8c8c8"
                btn.setAttribute("disabled", true)
            }
        }

        // 绑定事件
        btn.addEventListener("click", handleClick)
        room.addEventListener("input", validate)
        username.addEventListener("input", validate)
    }
)()