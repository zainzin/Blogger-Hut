<template>
  <div>
    <div class="flex-row justify-content-sm-center row">
      <BlogCard v-for="blog in blogs" :key="blog.Id" :blog="blog"></BlogCard>
    </div>
    <nav aria-label="Page navigation example">
      <ul class="pagination justify-content-center">
        <li :class="currentPage == 1 ? 'disabled' : ''" class="page-item disabled"><a class="page-link" href="#">Previous</a></li>
        <li v-for="page in totalBlogs" :class="currentPage == page ? 'active' : ''" class="page-item"><a class="page-link" href="#">{{page}}</a></li>
        <li :class="currentPage == totalBlogs ? 'disabled' : ''" class="page-item"><a class="page-link" href="#">Next</a></li>
      </ul>
    </nav>
  </div>
</template>

<script>
  import BlogCard from './BlogCard'

  export default {
    name: "Blogs",
    data() {
      return {
        blogs: [],
        totalBlogs: 0,
        currentPage: 1
      }
    },
    created() {
      fetch("http://localhost:3000/blogs")
        .then((resp) => resp.json())
        .then((data) => {
          this.blogs = data.Blogs
          this.totalBlogs = Math.ceil(data.TotalRows/10)
          this.currentPage = data.CurrentPage
        })
        .catch((err) => this.$noty.error(err))
    },
    components: {BlogCard}
  }
</script>

<style scoped>
nav {
  margin-top: 5%;
}
</style>
