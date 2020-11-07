<template>
  <div>
    <h2>How do you want to edit?</h2>
    <form>
      <input type="text" v-model="joke">
      <button v-on:click.prevent="updateJoke">Edit</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data: function() {
    return {
      joke:""
    }
  },
methods: {
     updateJoke() {
       {
        axios.put("http://localhost:8080/updateJoke",{
        id: this.$route.params.id,
        joke:this.joke
      })
      .then(response => {
        console.log(response)
        this.$router.push(`/jokes`)
      })
      }


    }
  }
}
</script>

<style>

</style>
