<template>
  <div>
    <FromEcho />
    <SearchJokes v-on:search-text="searchText" />
    <Joke v-for="joke in jokes" :key="joke.id" :id="joke.id" :joke="joke.joke" />
  </div>

</template>

<script>
import axios from 'axios';
import Joke from '../../components/Joke';
import SearchJokes from "../../components/SearchJokes";
import FromEcho from "../../components/FromEcho"

export default {
  components: {
    Joke,
    SearchJokes
  },
  data() {
    return {
      jokes: []
    }
  },
  async created() {
    const config = {
      headers: {
        'Accept': 'application/json'
      }
    }
    try {
      const res = await axios.get('http://localhost:8080/fetchJokes', config)
      console.log(res)
      this.jokes = res.data;
    } catch (err) {
      console.log(err)
    }

  },
  methods: {
    async searchText(text) {
      const config = {
      headers: {
        'Accept': 'application/json'
      }
    }
    try {
      const res = await axios.get(`https://icanhazdadjoke.com/search?term=${text}`, config);

      this.jokes = res.data.results;
    } catch (err) {
      console.log(err)
    }
   }
  },
  head() {
    return {
      title: 'Welcome to dad jokes',
      meta: [
        {
          hid: 'discription',
          name: 'description',
          content: "Best place for dad jokes"
        }
      ]
    }
  }
}
</script>
