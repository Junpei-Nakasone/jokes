<template>
  <nuxt-link :to="'jokes/' + id">
    <div class="joke">
    <p>{{ joke }}</p>
    <button >
      <nuxt-link :to="'update/' + id">Edit</nuxt-link>
    </button>
    <button v-on:click="deleteJoke(id)">
      Delete
    </button>
  </div>
  </nuxt-link>
</template>

<script>
import axios from 'axios'
export default {
  name: 'joke',
    props: ['joke', 'id'],
  data: {
    // name: 'joke',
    // props: ['joke', 'id']
  },
  methods: {
    deleteJoke() {
      // const params = {
      //   id: this.id
      // }
      alert(this.id)
      axios.delete("http://localhost:8080/deleteJoke",{ params:{id: this.id} })
      .then(res => {
        console.log(res)
        this.$router.push('/jokes')
      })
      .catch(err => {
        console.log(err)
      })
    }
  }

}
</script>

<style>
.joke {
  padding: 1rem;
  border: 1px dotted #ccc;
  margin: 1rem 0;
}
.updatebutton:hover {
  cursor: pointer;
  color:lemonchiffon;
}
</style>
