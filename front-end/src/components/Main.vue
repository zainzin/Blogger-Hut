<template>
  <div class="container">
    <h1>{{ msg }}</h1>
    <h2>{{ subtitle }}</h2>
    <div class="flex-row justify-content-sm-center row">
      <BlogCard v-for="blog in blogs" :key="blog.Id" :blog="blog"></BlogCard>
    </div>
  </div>

</template>

<script>
  import BlogCard from './BlogCard'
  export default {
    name: 'Main',
    data() {
      return {
        msg: 'Welcome to Your Bloggers Hut App',
        subtitle: 'Ten most recent blogs posted',
        blogs: []
      }
    },
    components: {BlogCard},
    created() {
      fetch(`http://localhost:3000/recent-blogs`)
        .then((resp) => resp.json())
        .then((json) => {
          this.blogs = json;
        })
        .catch((err) => {
          this.$noty.error("Oops, something went wrong!. " + err)
      })
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
