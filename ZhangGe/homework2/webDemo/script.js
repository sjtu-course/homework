let header1 = document.querySelector('h1');
let roomInput = document.querySelector('#room');
let userInput = document.querySelector('#user');
let loginButton = document.querySelector('button');
let stringTestFlag = false;

// Initial : disabled the button
loginButton.disabled = true;

function login() {
  //header1.textContent = '登录成功！';
  //roomInput.value = 'Hello';
  stringTest();
  if (stringTestFlag){
	alert('登录成功！');
  }
  else{
	alert('输入必须同时含有数字、字母和_！');
  }
}

// detect the instant change of the input
//https://developer.mozilla.org/zh-CN/docs/Web/API/GlobalEventHandlers/oninput
roomInput.oninput = handleRoomInput; 
function handleRoomInput(e) { 
if (roomInput.value !== "" && userInput.value !== ""){
	header1.textContent = '全部输入！';
	loginButton.disabled = false;
	loginButton.style.background = "#5286ed";
	
}
else{
	header1.textContent = '未全部输入！';
	loginButton.disabled = true;
}

}

userInput.oninput = handleUserInput; 
function handleUserInput(e) { 
if (roomInput.value !== "" && userInput.value !== ""){
	header1.textContent = '全部输入！';
	loginButton.disabled = false;
	loginButton.style.background = "#5286ed";
}
else{
	header1.textContent = '未全部输入！';
	loginButton.disabled = true;
	loginButton.style.background = "#c8c8c8";
}
}

//Test whether has letters, numbers and _ at the same time
function stringTest() {
	var roomHasLetter = false;
	var roomHasNumber = false;
	var roomHas_ = false;
	
	var userHasLetter = false;
	var userHasNumber = false;
	var userHas_ = false;
	
	var hasLetter = 0;
	var hasNumber = 0;
	var has_ = 0;
	var letterSet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
	var numberSet = '0123456789';
	
	/*Test whether has letters*/
	for (var i = 0; i < letterSet.length; i++){
		if (roomInput.value.indexOf(letterSet[i]) === -1) {
			roomHasLetter = true;
			
		}
		if (userInput.value.indexOf(letterSet[i]) === -1){
			userHasLetter = true;
		}
		
		if (roomHasLetter && userHasLetter){
			hasLetter = 1;
			break;
		}
	}
	
	/*Test whether has numbers*/
	for (var i = 0; i < numberSet.length; i++){
		if (roomInput.value.indexOf(numberSet[i]) === -1) {
			roomHasNumber = true;
			
		}
		if (userInput.value.indexOf(numberSet[i]) === -1){
			userHasNumber = true;
		}
		
		if (roomHasNumber && userHasNumber){
			hasNumber = 1;
			break;
		}
	}
	
	/*Test whether has _*/
	if (roomInput.value.indexOf('_') !== -1 || userInput.value.indexOf('_') !== -1){
		has_ = 1;
	}
	
	/*if has letters, numbers and _ at the same time*/
	if (hasLetter + hasNumber + has_ === 3){
		stringTestFlag = true;
	}

}

loginButton.onclick = function() {
	login();
}
