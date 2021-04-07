<template>
  <div class="Requester">
    <button @click="requestOpen">Request open</button>
    <button @click="requestClosed">Request closed</button>

    <div class="status">Result: {{ resultCode }}</div>
  </div>
</template>

<script>
export default {
  name: 'Requester',
  data: () => ({
    resultCode: null
  }),
  methods: {
    async requestOpen() {
      this.resultCode = await this.$store.dispatch('upstream/requestOpen')
    },
    async requestClosed() {
      this.resultCode = await this.$store.dispatch('upstream/requestClosed')
    }
  }
}
</script>

<style lang="css" scoped>
.Requester {
}

button {
  padding: 1em;

  border: 2px solid #41b883;
  border-radius: 4px;

  background: white;
}

button:not(:first-child) {
  margin-left: 1em;
}

.status {
  margin-top: 1em;
}
</style>
