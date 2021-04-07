<template>
  <TopBar />
  <img alt="Vue logo" src="./assets/logo.png" />
  <LoginBox />
  <br />
  <Requester />
</template>

<script>
import LoginBox from './components/LoginBox.vue'
import TopBar from './components/TopBar.vue'
import Requester from './components/Requester.vue'

export default {
  components: { Requester, TopBar, LoginBox }
}
</script>

<style>
html,
body {
  margin: 0;
  padding: 0;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
