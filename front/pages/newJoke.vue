<template>
  <form v-on:submit.prevent="newJoke">
    New Joke:
    <input type="text" v-model="joke" />
    <button type="submit">Add</button>
  </form>
</template>

<script>
import axios from "axios";

export default {
  data: function() {
    return {
      joke: ""
    };
  },
  created: function() {
    console.log("created run.");
  },
  methods: {
    newJoke() {
      axios
        .post("http://localhost:8080/addJoke", {
          joke: this.joke
        })
        .then(response => {
          console.log(response);
        });
      this.joke = "";
      this.$router.push(`/jokes`);
      console.log("first?");
    }
  }
};
</script>

<style>
</style>
