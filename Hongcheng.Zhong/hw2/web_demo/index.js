(
    function () {
        let room = document.getElementById("room")
        let user = document.getElementById("user")
        let btn = document.getElementById("submit")

        function handleClick() {
            let room_value = room.value
            let user_value = user.value
            let regexp = /^([a-zA-Z]|[0-9]|_)+$/
            if (regexp.test(room_value) && regexp.test(user_value)) {
                alert("成功!")
            } else {
                alert("失败!")
            }
        }

        function validate() {
            let room_value = room.value
            let user_value = user.value
            if (room_value && user_value) {
                btn.style.backgroundColor="#5286ed"
                btn.removeAttribute("disabled")
            } else {
                btn.style.backgroundColor="#c8c8c8"
                btn.setAttribute("disabled", true)
            }
        }

        btn.addEventListener("click", handleClick)
        room.addEventListener("input", validate)
        user.addEventListener("input", validate)
    }
)()
