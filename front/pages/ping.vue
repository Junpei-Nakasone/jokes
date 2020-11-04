<template>
  <div>
    <p>test</p>

    <h3>{{ msg }}</h3>
    </div>

</template>

<script>
import axios from 'axios';
export default {
  // async asyncData({ app }) {
  //   const res = await app.$axios.get('http://localhost:8080/ping');
  //   console.log(res)
  //   return {
  //     msg: res.data
  //   }
  // }
  async asyncData ({ app }) {
    return await app.$axios.get('http://localhost:8080/ping')
    .then((res) => {
      return {msg:res.data}
    })
    .catch((e) => {
      return {msg:"server is sleeping.."}
    })
  }
}
</script>

export default {
  data() {
    return {
      joke: {}
    }
  },
  async created() {
    const config = {
      headers: {
        'Accept': 'application/json'
      }
    }
    try {
      const res = await axios.get(`https://icanhazdadjoke.com/j/${this.$route.params.id}`, config);

      this.joke = res.data.joke;
    } catch (err) {
      console.log(er
