<template>
  <div>
    <div class="flex-row justify-content-sm-center row">
      <BlogCard v-for="blog in blogs" :key="blog.Id" :blog="blog"></BlogCard>
    </div>
    <nav aria-label="Page navigation example">
      <ul class="pagination justify-content-center">
        <li :class="currentPage == 1 ? 'disabled' : ''" class="page-item"><a class="page-link"
                                                                                      href="javascript:void(0)"
                                                                                      @click="prevPage()">Previous</a>
        </li>
        <li v-for="page in totalBlogs" :class="currentPage == page ? 'active' : ''" class="page-item"><a
          class="page-link" href="javascript:void(0)" @click="goToPage($event)">{{page}}</a></li>
        <li :class="currentPage == totalBlogs ? 'disabled' : ''" class="page-item"><a class="page-link"
                                                                                      href="javascript:void(0)"
                                                                                      @click="nextPage()">Next</a></li>
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
      this.fetchBlogsPerPage()
    },
    components: {BlogCard},
    methods: {
      goToPage(event) {
        const pageNumber = event.target.innerHTML;
        if (pageNumber != this.currentPage) {
          this.fetchBlogsPerPage(pageNumber)
        }
      },
      nextPage() {
        const nextPage = this.currentPage + 1
        if (nextPage <= this.totalBlogs) {
          this.fetchBlogsPerPage(nextPage)
        }
      },
      prevPage() {
        const prevPage = this.currentPage - 1
        if (prevPage > 0) {
          this.fetchBlogsPerPage(prevPage)
        }
      },
      fetchBlogsPerPage(page) {
        if (!page) {
          page = ''
        } else {
          page = `?skip=${page}`
        }
        fetch(`http://localhost:3000/blogs${page}`)
          .then((resp) => resp.json())
          .then((data) => {
            this.blogs = data.Blogs
            this.totalBlogs = Math.ceil(data.TotalRows / 10)
            this.currentPage = data.CurrentPage
            console.log(this.currentPage)
          })
          .catch((err) => this.$noty.error(err))
      }
    }
  }
</script>

<style scoped>
  nav {
    margin-top: 5%;
  }
</style>
