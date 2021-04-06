<template>
  <div class="LoginBox">
    <button class="primary" v-if="!isAuthenticated" @click="$store.dispatch('auth/login')">Login</button>
    <button v-else @click="$store.dispatch('auth/logout')">Logout</button>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'LoginBox',
  computed: {
    ...mapGetters('auth', ['isAuthenticated'])
  },
}
</script>

<style lang="css" scoped>
.LoginBox {
  display: flex;
  flex-direction: column;
  justify-content: center;

  width: 100%;
}

button {
  padding: 1em;
  margin: 1em;

  border: 2px solid #41b883;
  border-radius: 4px;

  font-weight: bold;
}

button.primary {
  color: #35495e;
  background: #41b883;
}
</style>
