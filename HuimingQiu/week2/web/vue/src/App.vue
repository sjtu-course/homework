<template>
  <div id="app">
    <Login/>
  </div>
</template>
<script>
import Login from './views/Login'
export default {
  name: 'App',
  components: {
    Login
  }
}
</script>

<style>
* {
  padding: 0;
  margin: 0;
}

body {
  background-color: #919191;
}

#app {
  display: flex;
  flex-direction: row;
  align-items: center;
  min-height: 100vh;
}
</style>
