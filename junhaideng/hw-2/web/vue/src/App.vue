<template>
  <div id="app">
    <WebRTC/>
  </div>
</template>

<script>
import WebRTC from './components/WebRTC.vue'

export default {
  name: 'App',
  components: {
    WebRTC
  }
}
</script>

<style>
#app {
  height: 100%;
}
html, body {
  height: 100%;
  width: 100%;
  padding: 0;
  margin: 0;
  background-color: #919191;
}

</style>
