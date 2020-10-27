<template>
  <div class="login-container">
    <h3 class="login-title">登录</h3>
    <input type="text" class="room-id" placeholder="房间ID" v-model="roomId" />
    <input type="text" class="user-id" placeholder="用户ID" v-model="userId" />
    <button
      :class="{ 'enter-button': true, 'active-button': buttonActive }"
      :disabled="!buttonActive"
      @click="onButtonClick"
    >
      进入房间
    </button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      roomId: "",
      userId: ""
    };
  },
  methods: {
    onButtonClick() {
      if (this.isValidString(this.roomId) && this.isValidString(this.userId)) {
        alert("登录成功");
      } else {
        alert("登录失败");
      }
    },
    isValidString(s) {
      return /[_a-zA-Z0-9]+/.test(s);
    }
  },
  computed: {
    buttonActive() {
      return this.roomId !== "" && this.userId !== "";
    }
  }
};
</script>

<style>
.login-container {
  background-color: white;
  width: 425px;
  height: 380px;
  margin: auto;
  padding: 45px 55px 0 50px;
  border-radius: 15px;
  display: flex;
  flex-direction: column;
}

.login-title {
  font-size: 30px;
  color: black;
}

.room-id {
  height: 35px;
  margin-top: 55px;
  border-radius: 2px;
  font-size: 16px;
  color: black;
  padding-left: 20px;
}

.user-id {
  height: 40px;
  margin-top: 25px;
  border-radius: 2px;
  font-size: 16px;
  color: black;
  padding-left: 20px;
}

.enter-button {
  margin-top: 40px;
  font-size: 18px;
  height: 50px;
  background-color: #c8c8c8;
  border-radius: 2px; /*加上感觉更好看*/
  color: white;
}

.active-button {
  background-color: #5286ed;
}

/* 修改placeholder的颜色 */
::-webkit-input-placeholder {
  color: #c1c1c1;
}
:-moz-placeholder {
  color: #c1c1c1;
}
::-moz-placeholder {
  color: #c1c1c1;
}
:-ms-input-placeholder {
  color: #c1c1c1;
}
</style>
