<template>
  <form v-on:prevent="newJoke">
    New Joke: <input type="text" v-model="joke">
    <button type="submit">Add</button>
  </form>
</template>

<script>
export default {
  data: function() {
    return {
      joke: ""
    }
  },
  created: function() {
    console.log("created run.")
  },
  methods: {
    async newJoke() {
      await this.$axios.post('http://localhost:8080/addJoke', {
        joke: this.joke
      })
      this.$router.push(`/jokes`)
    }
  }
}
</script>

<style>

</style>
