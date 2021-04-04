<template>
  <div class="TopBar">
    <span v-if="isAuthenticated">
      {{ user }}
    </span>
    <span v-else>
      Not logged in
    </span>
  </div>
</template>

<script>
import { mapGetters, mapState } from 'vuex'

export default {
  name: 'TopBar',
  computed: {
    ...mapGetters('auth', ['isAuthenticated']),
    ...mapState('auth', ['user']),
  },
  async created() {
    await this.$store.dispatch('auth/refresh')
  },
}
</script>

<style lang="css" scoped>
.TopBar {
  background: #41b883;

  padding: 1em;

  text-align: right;
}
</style>
