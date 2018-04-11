<template>
  <div>
    <h1>{{ msg }}</h1>
    <router-link to="/authors">Authors</router-link>
  </div>
</template>

<script>

export default {
  name: 'Main',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>

<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
