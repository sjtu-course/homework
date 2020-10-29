<template>
  <div class="container">
    <h1 class="text">登录</h1>
    <input type="text" placeholder="房间ID" v-model="rid" />
    <input type="text" placeholder="用户ID" v-model="uid" />
    <button :class="{ active: rid || uid }" @click="login()" :disabled="!(rid || uid)">进入房间</button>
  </div>
</template>

<script>
export default {
  name: "HelloWorld",
  data() {
    return {
      rid: "",
      uid: "",
    };
  },
  methods: {
    check(str) {
      console.log(str);
      let re = /^[0-9a-zA-Z_]*$/g;
      if (!re.test(str)) return false;
      else return true;
    },
    login() {
      if (this.check(this.uid) && this.check(this.rid)) alert("成功");
      else alert("失败")
    },
  },
};
</script>

<style scoped lang="scss">
.container {
  .text {
    display: flex;
  }
  height: 380px;
  position: absolute;
  top: 45px;
  left: 0;
  right: 0;
  margin: auto;
  width: 320px;
  display: flex;
  flex-direction: column;
  input {
    padding: 0 10px;
    height: 35px;
    border: 1px solid gray;
    border-radius: 2px;
    &:last-of-type {
      margin-bottom: 40px;
    }
    &:first-of-type {
      margin-bottom: 25px;
    }
  }
  button {
    padding: 10px;
    font-size: 18px;
    color: white;
    background: #c8c8c8;
    outline: none;
    border: none;
    border-radius: 2px;
    
  }
  .active {
    background: #5286ed;
  }
}
</style>
