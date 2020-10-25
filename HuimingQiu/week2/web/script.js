const roomIdInput = document.querySelector(".room-id");
const userIdInput = document.querySelector(".user-id");
const enterButton = document.querySelector(".enter-button");

let roomId = "";
let userId = "";

const changeButton = () => {
  if (roomId !== "" && userId !== "") {
    enterButton.style.backgroundColor = "#5286ed";
    enterButton.disabled = false;
  } else {
    enterButton.style.backgroundColor = "#c8c8c8";
    enterButton.disabled = true;
  }
};

// 使用正则表达式匹配字符串
const isValidString = (s) => {
  return /[_a-zA-Z0-9]+/.test(s);
};

roomIdInput.addEventListener("input", (e) => {
  roomId = e.target.value;
  changeButton();
});
userIdInput.addEventListener("input", (e) => {
  userId = e.target.value;
  changeButton();
});
enterButton.addEventListener("click", () => {
  if (isValidString(roomId) && isValidString(userId)) {
    alert("登录成功");
  } else {
    alert("登录失败");
  }
});
