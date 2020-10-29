<template>
    <div class="container">
      <div class="content">
        <div class="title">
          登录
        </div>
        <div class="form">
          <div>
            <input  v-model="room" @input="validate" class="input room" type="text" placeholder="房间ID"/>
          </div>
          <div>
            <input v-model="username" @input="validate" class="input username" type="text" placeholder="用户ID"/>
          </div>
          <div>
            <button @click="handleClick" ref="btn" class="btn default" :disabled="disabled">进入房间</button>
          </div>
        </div>
      </div>
    </div>
</template>

<script>
export default {
  name: 'WebRTC',
  data(){
    return{
      room: "",
      username: "",
      disabled: true,
    }
  },
  methods:{
    handleClick(){
      let regexp = /^([a-zA-Z]|[0-9]|_)+$/
      if(regexp.test(this.room) && regexp.test(this.username)){
        alert("成功")
      }else{
        alert("失败")
      }
    },
    validate(){
      if(this.room && this.username){
        this.$refs["btn"].className = "btn success"
        this.disabled = false
      }else{
        this.$refs["btn"].className = "btn default"
        this.disabled = true
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.container{
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.content {
  height: 380px;
  width: 425px;
  background-color: white;
  border-radius: 15px;
}

.title {
  position: relative;
  margin-top: 45px;
  margin-left: 50px;
  font-size: 30px;
  color: black;
  font-weight: bold;
  letter-spacing: 3px;
}

.form {
  margin-left: 50px;
  margin-top: 40px;
}

.room {
  height: 35px;
}

.username {
  height: 40px;
  margin-top: 25px;
}

.input {
  padding-left: 30px;
  font-size: 16px;

  /* 因为要设置padding，所以实际上width!=320 */
  width: 286px;
  border: 1px solid #c1c1c1;
  border-radius: 2px;
  outline: 0;
}

/* 更改placeholder里面的样式 */
input::-ms-input-placeholder {
  color: #c1c1c1;
}

input::-webkit-input-placeholder {
  color: #c1c1c1;
}

.btn {
  color: white;
  width: 320px;
  font-size: 18px;
  letter-spacing: 4px;
  margin-top: 40px;
  height: 40px;
  border-radius: 2px;
  border-width: 0;
  outline: 0;
}


.default{
  background-color: #c8c8c8;
  cursor: not-allowed;
}

.success {
  background-color: #5286ed;
  cursor: pointer;
}
</style>
